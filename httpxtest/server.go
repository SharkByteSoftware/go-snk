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

type ServerBuilder struct {
	t              *testing.T
	defaultHandler http.HandlerFunc
	routes         map[string]http.HandlerFunc
	options        []Option
}

func NewServerBuilder(t *testing.T, options ...Option) *ServerBuilder {
	return &ServerBuilder{
		t:              t,
		defaultHandler: defaultHandler,
		routes:         make(map[string]http.HandlerFunc),
		options:        options,
	}
}

func (sb *ServerBuilder) Build() *httptest.Server {
	ts := httptest.NewServer(http.HandlerFunc(sb.handler))
	sb.t.Cleanup(ts.Close)

	return ts
}

func (sb *ServerBuilder) BuildTLS() *httptest.Server {
	ts := httptest.NewTLSServer(http.HandlerFunc(sb.handler))
	sb.t.Cleanup(ts.Close)

	return ts
}

func (sb *ServerBuilder) On(statusCode int, response any, options ...Option) *ServerBuilder {
	sb.t.Helper()

	return sb.OnFunc(func(w http.ResponseWriter, _ *http.Request) {
		writeResponse(w, statusCode, response)
	})
}

func (sb *ServerBuilder) OnFunc(handler http.HandlerFunc) *ServerBuilder {
	sb.t.Helper()
	sb.defaultHandler = handler

	return sb
}

func (sb *ServerBuilder) OnRoute(method string, route string, statusCode int, response any, options ...Option) *ServerBuilder {
	sb.t.Helper()

	return sb.OnRouteFunc(method, route,
		func(w http.ResponseWriter, _ *http.Request) { writeResponse(w, statusCode, response) }, options...,)
}

func (sb *ServerBuilder) OnRouteFunc(method string, route string, handler http.HandlerFunc, options ...Option) *ServerBuilder {
	sb.t.Helper()

	if helpers.IsEmpty(method) {
		sb.t.Fatalf("method cannot be empty")
		return sb
	}

	key := routeKey(method, route)

	_, exists := sb.routes[key]
	if exists {
		sb.t.Fatalf("handler already defined for: %q", key)
	}

	sb.routes[key] = func(w http.ResponseWriter, req *http.Request) {
		slicex.Apply(options, func(option Option) { option(w, req) })
		handler(w, req)
	}

	return sb
}

func (sb *ServerBuilder) handler(w http.ResponseWriter, req *http.Request) {
	key := routeKey(req.Method, req.URL.Path)
	handler, ok := sb.routes[key]

	slicex.Apply(sb.options, func(option Option) { option(w, req) })

	if ok {
		handler(w, req)

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
		w.WriteHeader(statusCode)
	case string:
		w.WriteHeader(statusCode)
		_, _ = io.WriteString(w, value)
	case []byte:
		w.WriteHeader(statusCode)
		_, _ = w.Write(value)
	case json.RawMessage:
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(statusCode)
		_, _ = w.Write(value)
	default:
		w.Header().Set("Content-Type", "application/json")
		bytes, err := jsonx.EncodeBytes(value)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(statusCode)
		_, _ = w.Write(bytes)
	}
}

func routeKey(method, route string) string {
	return strings.ToUpper(method) + " " + strings.ToLower(route)
}

// Feature suggestions
// The commented-out block at the bottom of the test points at the most valuable missing feature — request recording/assertions:
// 1.
// Request recorder. Have Build return a small wrapper that embeds *httptest.Server (keeps .URL working) and records each inbound
// request (method, path, query, headers, captured body). Then:
// ts.Requests()              // []RecordedRequest
// ts.RequestsFor("/v1/horton")
// ts.Called("/v1/horton")    // bool
// ts.CallCount("/v1/horton") // int
// Capture the body eagerly (read + restore) so assertions don't fight a drained reader.
// 2.
// Sequenced / queued responses for retry and pagination tests — successive calls to the same route return different responses:
// .OnRouteGetSeq("/v1/x", resp1, resp2, resp3)
// 3.
// Latency / delay injection to exercise client timeouts and context cancellation:
// .WithDelay(route, 200*time.Millisecond)
// 4.
// Request matchers / expectations — match on header, query param, or body, and optionally t.Error if an expected request
// never arrives (ts.AssertExpectations(t)), mock-style.
// 5.
// TLS variant — BuildTLS(t) using httptest.NewTLSServer for testing clients against HTTPS / custom cert pools.
// 6.
// Response headers — a way to set arbitrary headers per route (e.g. Retry-After, Location, ETag), since header behavior
// is a common thing to test.
// 7.
// Handler escape hatch per route — OnRouteFunc(method, route, http.HandlerFunc) for cases the declarative API doesn't cover
// (streaming, chunked, custom status logic). You have WithDefaultHandler; a per-route equivalent rounds it out.
