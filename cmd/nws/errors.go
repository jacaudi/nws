package nws

import (
	"errors"
	"fmt"
)

// Sentinel errors returned by Client construction and methods.
var (
	// ErrUserAgentRequired is returned when a request is attempted with an
	// empty UserAgent. api.weather.gov requires a User-Agent header.
	ErrUserAgentRequired = errors.New("nws: User-Agent is required")

	// ErrInvalidUnits is returned by WithUnits when the provided value is
	// not "us", "si", or empty.
	ErrInvalidUnits = errors.New(`nws: units must be "us", "si", or empty`)

	// ErrInvalidBaseURL is returned by WithBaseURL when the value is not a
	// parseable absolute URL with an http or https scheme.
	ErrInvalidBaseURL = errors.New("nws: BaseURL must be an absolute http(s) URL")
)

// APIError represents a non-2xx response from api.weather.gov.
//
// Consumers can use errors.As(err, &apiErr) to inspect the StatusCode and
// distinguish, e.g., a 404 (unknown station) from a 429 (rate limited).
type APIError struct {
	StatusCode int    // HTTP status code (e.g. 404)
	Status     string // HTTP status text (e.g. "404 Not Found")
	URL        string // URL that produced the error
	Body       string // Response body, truncated to 1 KiB
}

// Error implements the error interface.
func (e *APIError) Error() string {
	if e.Body == "" {
		return fmt.Sprintf("nws: %s for %s", e.Status, e.URL)
	}
	return fmt.Sprintf("nws: %s for %s: %s", e.Status, e.URL, e.Body)
}
