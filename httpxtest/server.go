package httpxtest

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/SharkByteSoftware/go-snk/conditional"
	"github.com/SharkByteSoftware/go-snk/jsonx"
)

type ServerBuilder struct {
	defaultHandler http.HandlerFunc
}

func NewServerBuilder() *ServerBuilder {
	return &ServerBuilder{
		defaultHandler: defaultHandler,
	}
}

func (sb *ServerBuilder) Build(t testing.TB) *httptest.Server {
	t.Helper()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		sb.defaultHandler(w, req)
	}))

	t.Cleanup(ts.Close)

	return ts
}

func (sb *ServerBuilder) WithDefaultHandler(handler http.HandlerFunc) *ServerBuilder {
	sb.defaultHandler = handler

	return sb
}

func (sb *ServerBuilder) On(statusCode int, result any) *ServerBuilder {

	sb.defaultHandler = func(w http.ResponseWriter, r *http.Request) {
		err := jsonx.Encode(w, result)
		w.WriteHeader(conditional.If(err != nil, http.StatusBadRequest, statusCode))
	}

	return sb
}

func (sb *ServerBuilder) OnStr(statusCode int, result any) *ServerBuilder {
	return sb
}

func (sb *ServerBuilder) OnGet(route string, statusCode int, result any) *ServerBuilder {
	return sb
}

func (sb *ServerBuilder) OnGetStr(route string, statusCode int, result string) *ServerBuilder {
	return sb
}

func defaultHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
	_, _ = w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
}
