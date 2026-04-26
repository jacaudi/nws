package nws

import (
	"fmt"
	"net/http"
	"net/url"
	"runtime/debug"
	"strings"
	"time"
)

// Default values for a freshly constructed Client.
const (
	defaultBaseURL = "https://api.weather.gov"
	defaultAccept  = "application/ld+json"
	defaultTimeout = 30 * time.Second
)

// version is the library version, read from build info at package init.
// Falls back to "dev" when build info is unavailable (e.g. `go run`).
var version = func() string {
	if info, ok := debug.ReadBuildInfo(); ok {
		if v := info.Main.Version; v != "" && v != "(devel)" {
			return v
		}
	}
	return "dev"
}()

// defaultUserAgent is the User-Agent header sent when none is specified.
var defaultUserAgent = fmt.Sprintf("nws-go/%s (+https://github.com/jacaudi/nws)", version)

// Client is the entry point for all NWS API calls.
//
// Construct a Client with NewClient (returns an error if any option is
// invalid) or MustNewClient (panics on error). A Client is safe for
// concurrent use; do not mutate its fields after the first request.
type Client struct {
	BaseURL    string
	UserAgent  string
	Accept     string
	Units      string
	HTTPClient *http.Client
}

// Option configures a *Client. Options return an error so they can validate
// their inputs at construction time.
type Option func(*Client) error

// NewClient returns a Client with sensible defaults. Pass options to
// override individual fields. With no options NewClient cannot fail.
//
// On error the first failing option is returned; subsequent options are
// not applied.
func NewClient(opts ...Option) (*Client, error) {
	c := &Client{
		BaseURL:   defaultBaseURL,
		UserAgent: defaultUserAgent,
		Accept:    defaultAccept,
		Units:     "",
		HTTPClient: &http.Client{
			Timeout: defaultTimeout,
		},
	}
	for _, opt := range opts {
		if err := opt(c); err != nil {
			return nil, err
		}
	}
	return c, nil
}

// MustNewClient is NewClient that panics on error. Useful when a
// misconfigured client is a programmer error and there is no graceful
// recovery, e.g. var declarations or main().
func MustNewClient(opts ...Option) *Client {
	c, err := NewClient(opts...)
	if err != nil {
		panic(err)
	}
	return c
}

// WithBaseURL overrides the API base URL. The value must be a parseable
// absolute URL with an http or https scheme. Trailing slashes are
// stripped, so callers do not need to worry about double-slash paths.
func WithBaseURL(rawURL string) Option {
	return func(c *Client) error {
		u, err := url.Parse(rawURL)
		if err != nil || !u.IsAbs() || (u.Scheme != "http" && u.Scheme != "https") {
			return ErrInvalidBaseURL
		}
		c.BaseURL = strings.TrimRight(rawURL, "/")
		return nil
	}
}

// WithUserAgent overrides the User-Agent header sent with every request.
// Empty values are rejected with ErrUserAgentRequired.
func WithUserAgent(ua string) Option {
	return func(c *Client) error {
		if ua == "" {
			return ErrUserAgentRequired
		}
		c.UserAgent = ua
		return nil
	}
}

// WithAccept overrides the Accept header. Empty values are rejected
// with ErrAcceptRequired; pass the default explicitly if you need it.
//
// The library's response struct types are shaped for
// "application/ld+json"; consumers passing other values must handle
// the response shape themselves.
func WithAccept(accept string) Option {
	return func(c *Client) error {
		if accept == "" {
			return ErrAcceptRequired
		}
		c.Accept = accept
		return nil
	}
}

// WithUnits sets the unit system. Valid values are "us", "si", or empty.
// Any other value returns ErrInvalidUnits.
func WithUnits(units string) Option {
	return func(c *Client) error {
		switch units {
		case "us", "si", "":
			c.Units = units
			return nil
		default:
			return ErrInvalidUnits
		}
	}
}

// WithHTTPClient overrides the underlying *http.Client. A nil value is
// rejected; pass http.DefaultClient explicitly if that is desired.
func WithHTTPClient(hc *http.Client) Option {
	return func(c *Client) error {
		if hc == nil {
			return fmt.Errorf("nws: WithHTTPClient: client must not be nil")
		}
		c.HTTPClient = hc
		return nil
	}
}

// DefaultClient is used by the package-level wrapper functions
// (GetPoints, RadarStation, etc.). Replace it before making calls to
// customize default behavior. Mutation after first use is not
// goroutine-safe; build per-request *Client instances if you need that.
var DefaultClient = MustNewClient()
