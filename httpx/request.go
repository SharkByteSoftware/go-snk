package httpx

import (
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"sync"

	"github.com/SharkByteSoftware/go-snk/conditional"
)

// DoRawRequest sends an HTTP request with the given method and body, without any shared request logic.
func DoRawRequest(ctx context.Context, method, url string, body io.Reader, options ...Option) (*http.Response, error) {
	config, err := applyOptions(options)
	if err != nil {
		return nil, fmt.Errorf("DoRawRequest: %w", err)
	}

	return doRawRequest(ctx, method, url, body, config)
}

// DoRequest sends an HTTP request with the given method and body, applying shared request logic.
func DoRequest[T any](ctx context.Context, method, url string, body io.Reader, options ...Option) (*Response[T], error) {
	config, err := applyOptions(options)
	if err != nil {
		return nil, fmt.Errorf("DoRequest: %w", err)
	}

	resp, err := doRawRequest(ctx, method, url, body, config)
	if err != nil {
		return nil, err
	}

	defer func() { _ = resp.Body.Close() }()

	return DecodeResponse[T](resp, config)
}

func doRawRequest(ctx context.Context, method, url string, body io.Reader, config *ConfigOptions) (*http.Response, error) {
	if ctx == nil {
		return nil, NewTransportError(fmt.Errorf("%w: nil context", ErrOptions))
	}

	req, err := newRequestWithAppliedConfig(ctx, method, url, body, config)
	if err != nil {
		return nil, NewTransportError(err)
	}

	client := clientWithAppliedConfig(config)

	resp, err := client.Do(req)
	if err != nil {
		return nil, NewTransportError(err)
	}

	return resp, nil
}

// insecureTransport is shared across requests so connections are pooled and reused.
//nolint:gochecknoglobals
var insecureTransport = sync.OnceValue(func() *http.Transport {
	return newInsecureTransport(http.DefaultTransport)
})

// newInsecureTransport returns a transport with certificate verification disabled.
// When base is an *http.Transport it is cloned, preserving proxy-from-environment, HTTP/2,
// and dial/TLS handshake timeouts; otherwise a minimal transport with proxy support is used.
// No insecure settings are applied to base. Note that http.Transport.Clone initializes
// base's lazy HTTP/2 defaults as a side effect.
func newInsecureTransport(base http.RoundTripper) *http.Transport {
	typed, isTransport := base.(*http.Transport)

	// typed.Clone is a method value on a possibly-nil pointer; that is safe because
	// IfCallReturn only invokes the branch selected by isTransport.
	transport := conditional.IfCallReturn(isTransport,
		typed.Clone,
		//nolint:exhaustruct_v5
		func() *http.Transport { return &http.Transport{Proxy: http.ProxyFromEnvironment} },
	)

	// Clone any existing TLS config rather than replacing it, so settings such as the
	// NextProtos that enable HTTP/2 are kept. (*tls.Config).Clone is nil-safe.
	tlsConfig := conditional.IfCallReturn(transport.TLSClientConfig != nil,
		transport.TLSClientConfig.Clone,
		//nolint:exhaustruct_v5
		func() *tls.Config { return &tls.Config{} },
	)

	tlsConfig.InsecureSkipVerify = true
	tlsConfig.MinVersion = max(tlsConfig.MinVersion, tls.VersionTLS12)
	transport.TLSClientConfig = tlsConfig

	return transport
}

func clientWithAppliedConfig(config *ConfigOptions) *http.Client {
	if config.httpClient != nil {
		return config.httpClient
	}

	//nolint:exhaustruct_v5
	client := &http.Client{Timeout: config.timeout}

	if config.insecureSkipVerify {
		client.Transport = insecureTransport()
	}

	return client
}

func newRequestWithAppliedConfig(
	ctx context.Context,
	method string,
	baseURL string,
	body io.Reader,
	config *ConfigOptions,
) (*http.Request, error) {
	base, err := config.parseURLFunc(baseURL)
	if err != nil {
		return nil, err
	}

	if len(config.params) > 0 {
		query := base.Query()
		for key, values := range config.params {
			//nolint:modernize
			query[key] = values
		}

		base.RawQuery = query.Encode()
	}

	req, err := http.NewRequestWithContext(ctx, method, base.String(), body)
	if err != nil {
		return nil, fmt.Errorf("new request: %w", err)
	}

	if config.headers != nil {
		req.Header = config.headers.Clone()
	}

	return req, nil
}
