// Package httpxtest provides a set of utilities for testing HTTP servers.
package httpxtest

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/SharkByteSoftware/go-snk/helpers"
	"github.com/SharkByteSoftware/go-snk/jsonx"
	"github.com/SharkByteSoftware/go-snk/slicex"
)

// ServerBuilder is a builder for HTTP servers.
type ServerBuilder struct {
	t         *testing.T
	options   []Option
	onHandler http.HandlerFunc
	routes    map[string]http.HandlerFunc
}

// NewServerBuilder creates a new ServerBuilder.
func NewServerBuilder(t *testing.T, options ...Option) *ServerBuilder {
	t.Helper()

	return &ServerBuilder{
		t:         t,
		routes:    make(map[string]http.HandlerFunc),
		options:   options,
		onHandler: nil,
	}
}

// Build creates a new HTTP server.
func (sb *ServerBuilder) Build() *httptest.Server {
	ts := httptest.NewServer(http.HandlerFunc(sb.handler))
	sb.t.Cleanup(ts.Close)

	return ts
}

// BuildTLS creates a new HTTPS server.
func (sb *ServerBuilder) BuildTLS() *httptest.Server {
	ts := httptest.NewTLSServer(http.HandlerFunc(sb.handler))
	sb.t.Cleanup(ts.Close)

	return ts
}

// On defines a wildcard path handler.
// If the response is nil, the server always responds with 204 No Content regardless of statusCode.
func (sb *ServerBuilder) On(statusCode int, response any, options ...Option) *ServerBuilder {
	return sb.OnFunc(func(w http.ResponseWriter, _ *http.Request) {
		writeResponse(w, statusCode, response)
	}, options...)
}

// OnFunc defines a wildcard handler function.
func (sb *ServerBuilder) OnFunc(handler http.HandlerFunc, options ...Option) *ServerBuilder {
	if sb.onHandler != nil {
		panic("an On handler already defined")
	}

	sb.onHandler = func(w http.ResponseWriter, r *http.Request) {
		slicex.Apply(options, func(option Option) { option(w, r) })
		handler(w, r)
	}

	return sb
}

// OnSequence registers an ordered sequence of responses for the default handler.
func (sb *ServerBuilder) OnSequence(exhaust ExhaustBehavior, responses ...SequencedResponse) *ServerBuilder {
	if sb.onHandler != nil {
		panic("an On handler already defined")
	}

	sb.onHandler = sb.createSequenceFunc(exhaust, responses...)

	return sb
}


// OnRoute defines a handler for a specific route.
// If the response is nil, the server always responds with 204 No Content regardless of statusCode.
func (sb *ServerBuilder) OnRoute(method string, route string, statusCode int, response any, options ...Option) *ServerBuilder {
	return sb.OnRouteFunc(method, route,
		func(w http.ResponseWriter, _ *http.Request) { writeResponse(w, statusCode, response) }, options...)
}

// OnRouteFunc defines a handler for a specific route.
func (sb *ServerBuilder) OnRouteFunc(method string, route string, handler http.HandlerFunc, options ...Option) *ServerBuilder {
	sb.addRouteFunc(method, route, func(w http.ResponseWriter, req *http.Request) {
		slicex.Apply(options, func(option Option) { option(w, req) })
		handler(w, req)
	})

	return sb
}

// OnRouteSequence registers an ordered sequence of responses for a method/route.
func (sb *ServerBuilder) OnRouteSequence(method, route string, exhaust ExhaustBehavior, responses ...SequencedResponse) *ServerBuilder {
	sb.addRouteFunc(method, route, sb.createSequenceFunc(exhaust, responses...))
	return sb
}

func (sb *ServerBuilder) addRouteFunc(method, route string, handlerFunc http.HandlerFunc) {
	if helpers.IsEmpty(method) {
		panic("method cannot be empty")
	}

	key := routeKey(method, route)

	_, exists := sb.routes[key]
	if exists {
		panic("handler already defined for: " + key)
	}

	sb.routes[key] = handlerFunc
}

func (sb *ServerBuilder) createSequenceFunc(exhaust ExhaustBehavior, responses ...SequencedResponse) http.HandlerFunc {
	//nolint:exhaustruct
	entry := &routeEntry{exhaust: exhaust}

	for _, r := range responses {
		resp := r

		entry.handlers = append(entry.handlers, func(w http.ResponseWriter, req *http.Request) {
			slicex.Apply(resp.options, func(o Option) { o(w, req) })
			writeResponse(w, resp.statusCode, resp.body)
		})
	}

	return func(w http.ResponseWriter, r *http.Request) {
		entry.next()(w, r)
	}
}

func (sb *ServerBuilder) handler(w http.ResponseWriter, req *http.Request) {
	key := routeKey(req.Method, req.URL.Path)
	handler, ok := sb.routes[key]

	slicex.Apply(sb.options, func(option Option) { option(w, req) })

	if ok {
		handler(w, req)

		return
	}

	if sb.onHandler != nil {
		sb.onHandler(w, req)
	}

	defaultHandler(w, req)
}

func defaultHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	_, _ = w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
}

func writeResponse(w http.ResponseWriter, statusCode int, response any) {
	switch value := response.(type) {
	case nil:
		w.WriteHeader(http.StatusNoContent)
	case string:
		w.WriteHeader(statusCode)
		_, _ = io.WriteString(w, value)
	case []byte:
		w.WriteHeader(statusCode)
		_, _ = w.Write(value)
	case json.RawMessage:
		w.WriteHeader(statusCode)
		_, _ = w.Write(value)
	default:
		bytes, err := jsonx.EncodeBytes(value)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(statusCode)
		_, _ = w.Write(bytes)
	}
}

func routeKey(method, route string) string {
	return strings.ToUpper(method) + " " + route
}
