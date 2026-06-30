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
	t               *testing.T
	defaultHandler  http.HandlerFunc
	defaultSequence *routeEntry
	routes          map[string]http.HandlerFunc
	routes2         map[string]*routeEntry
	options         []Option
}

// NewServerBuilder creates a new ServerBuilder.
func NewServerBuilder(t *testing.T, options ...Option) *ServerBuilder {
	t.Helper()

	return &ServerBuilder{
		t:              t,
		defaultHandler: defaultHandler,
		routes:         make(map[string]http.HandlerFunc),
		options:        options,
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

// On defines a handler for the default route.
// If the response is nil, the server always responds with 204 No Content regardless of statusCode.
func (sb *ServerBuilder) On(statusCode int, response any, options ...Option) *ServerBuilder {
	return sb.OnFunc(func(w http.ResponseWriter, _ *http.Request) {
		writeResponse(w, statusCode, response)
	}, options...)
}

// OnFunc defines a handler for the default route.
func (sb *ServerBuilder) OnFunc(handler http.HandlerFunc, options ...Option) *ServerBuilder {
	key := "OnFunc"

	_, exists := sb.routes[key]
	if exists {
		panic("handler already defined for: " + key)
	}

	sb.routes[key] = func(w http.ResponseWriter, req *http.Request) {
		slicex.Apply(options, func(option Option) { option(w, req) })
		handler(w, req)
	}

	return sb
}

// OnRoute defines a handler for a specific route.
// If the response is nil, the server always responds with 204 No Content regardless of statusCode.
func (sb *ServerBuilder) OnRoute(method, route string, statusCode int, response any, options ...Option) *ServerBuilder {
	return sb.OnRouteSequence(method, route, ExhaustCycle,
		Response(statusCode, response, options...))
}

// OnRouteFunc registers a custom handler for a method/route.
func (sb *ServerBuilder) OnRouteFunc(method, route string, handler http.HandlerFunc, options ...Option) *ServerBuilder {
	wrapped := func(w http.ResponseWriter, r *http.Request) {
		slicex.Apply(options, func(o Option) { o(w, r) })
		handler(w, r)
	}
	sb.registerRoute(method, route, ExhaustCycle, []http.HandlerFunc{wrapped})

	return sb
}

// OnRouteSequence registers an ordered sequence of responses for a method/route.
func (sb *ServerBuilder) OnRouteSequence(method, route string, exhaust ExhaustBehavior, responses ...SequencedResponse) *ServerBuilder {
	handlers := slicex.Map(responses, func(r SequencedResponse) http.HandlerFunc {
		return func(w http.ResponseWriter, req *http.Request) {
			slicex.Apply(r.options, func(o Option) { o(w, req) })
			writeResponse(w, r.statusCode, r.body)
		}
	})
	sb.registerRoute(method, route, exhaust, handlers)
	return sb
}

// OnSequence registers an ordered sequence of responses for the default handler.
func (sb *ServerBuilder) OnSequence(exhaust ExhaustBehavior, responses ...SequencedResponse) *ServerBuilder {
	entry := &routeEntry{exhaust: exhaust}

	for _, r := range responses {
		resp := r
		entry.handlers = append(entry.handlers, func(w http.ResponseWriter, req *http.Request) {
			slicex.Apply(resp.options, func(o Option) { o(w, req) })
			writeResponse(w, resp.statusCode, resp.body)
		})
	}

	sb.defaultSequence = entry

	return sb
}

func (sb *ServerBuilder) registerRoute(method, route string, exhaust ExhaustBehavior, handlers []http.HandlerFunc) {
	if helpers.IsEmpty(method) {
		panic("method cannot be empty")
	}

	key := routeKey(method, route)
	if _, exists := sb.routes[key]; exists {
		panic("handler already defined for: " + key)
	}

	sb.routes2[key] = &routeEntry{handlers: handlers, exhaust: exhaust}
}

func (sb *ServerBuilder) handler(w http.ResponseWriter, req *http.Request) {
	slicex.Apply(sb.options, func(o Option) { o(w, req) })

	key := routeKey(req.Method, req.URL.Path)
	if entry, ok := sb.routes2[key]; ok {
		entry.next()(w, req)
		return
	}

	if sb.defaultSequence != nil {
		sb.defaultSequence.next()(w, req)
		return
	}

	sb.defaultHandler(w, req)
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
