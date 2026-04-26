package nws

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// errorBodyLimit is the maximum number of bytes from a non-2xx response
// body included in *APIError. NWS error responses are typically <500
// bytes; 1 KiB gives 2x headroom and keeps errors small enough to log.
const errorBodyLimit = 1024

// get performs a GET against the client's BaseURL+path, decodes a 2xx
// JSON response into out, and returns *APIError for non-2xx responses.
//
// The provided context controls cancellation and deadlines; pass
// context.Background() if no scope is needed.
func (c *Client) get(ctx context.Context, path string, out any) error {
	if c.UserAgent == "" {
		return ErrUserAgentRequired
	}

	reqURL := c.BaseURL + path
	if c.Units != "" {
		reqURL = reqURL + "?units=" + url.QueryEscape(c.Units)
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, reqURL, nil)
	if err != nil {
		return fmt.Errorf("nws: build request: %w", err)
	}
	req.Header.Set("User-Agent", c.UserAgent)
	req.Header.Set("Accept", c.Accept)

	httpClient := c.HTTPClient
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("nws: do request: %w", err)
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		body, _ := io.ReadAll(io.LimitReader(resp.Body, errorBodyLimit))
		return &APIError{
			StatusCode: resp.StatusCode,
			Status:     resp.Status,
			URL:        reqURL,
			Body:       string(body),
		}
	}

	if err := json.NewDecoder(resp.Body).Decode(out); err != nil {
		return fmt.Errorf("nws: decode response: %w", err)
	}
	return nil
}
