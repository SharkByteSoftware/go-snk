package httpx

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func Post[T any, R any](ctx context.Context, url string, headers http.Header, payload T,
	timeout time.Duration) (*R, error) {
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	req.Header = headers
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: timeout}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		repBody, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf(" %d: %s", resp.StatusCode, string(repBody))
	}

	var result R

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
