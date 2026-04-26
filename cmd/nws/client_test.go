package nws

import (
	"errors"
	"net/http"
	"strings"
	"testing"
	"time"
)

func TestNewClient_Defaults(t *testing.T) {
	c, err := NewClient()
	if err != nil {
		t.Fatalf("NewClient() error = %v, want nil", err)
	}
	if c.BaseURL != "https://api.weather.gov" {
		t.Errorf("BaseURL = %q, want https://api.weather.gov", c.BaseURL)
	}
	if c.Accept != "application/ld+json" {
		t.Errorf("Accept = %q, want application/ld+json", c.Accept)
	}
	if !strings.HasPrefix(c.UserAgent, "nws-go/") {
		t.Errorf("UserAgent = %q, want prefix nws-go/", c.UserAgent)
	}
	if c.Units != "" {
		t.Errorf("Units = %q, want empty", c.Units)
	}
	if c.HTTPClient == nil {
		t.Fatal("HTTPClient must not be nil by default")
	}
	if c.HTTPClient.Timeout != 30*time.Second {
		t.Errorf("HTTPClient.Timeout = %v, want 30s", c.HTTPClient.Timeout)
	}
}

func TestWithBaseURL(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		c, err := NewClient(WithBaseURL("https://example.com"))
		if err != nil {
			t.Fatalf("error = %v", err)
		}
		if c.BaseURL != "https://example.com" {
			t.Errorf("BaseURL = %q", c.BaseURL)
		}
	})
	t.Run("invalid", func(t *testing.T) {
		_, err := NewClient(WithBaseURL("not a url"))
		if !errors.Is(err, ErrInvalidBaseURL) {
			t.Errorf("err = %v, want ErrInvalidBaseURL", err)
		}
	})
	t.Run("relative", func(t *testing.T) {
		_, err := NewClient(WithBaseURL("/relative"))
		if !errors.Is(err, ErrInvalidBaseURL) {
			t.Errorf("err = %v, want ErrInvalidBaseURL", err)
		}
	})
}

func TestWithUserAgent(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		c, err := NewClient(WithUserAgent("myapp/1.0"))
		if err != nil {
			t.Fatalf("error = %v", err)
		}
		if c.UserAgent != "myapp/1.0" {
			t.Errorf("UserAgent = %q", c.UserAgent)
		}
	})
	t.Run("empty", func(t *testing.T) {
		_, err := NewClient(WithUserAgent(""))
		if !errors.Is(err, ErrUserAgentRequired) {
			t.Errorf("err = %v, want ErrUserAgentRequired", err)
		}
	})
}

func TestWithUnits(t *testing.T) {
	for _, u := range []string{"us", "si", ""} {
		c, err := NewClient(WithUnits(u))
		if err != nil {
			t.Errorf("WithUnits(%q): %v", u, err)
			continue
		}
		if c.Units != u {
			t.Errorf("Units = %q, want %q", c.Units, u)
		}
	}
	_, err := NewClient(WithUnits("invalid"))
	if !errors.Is(err, ErrInvalidUnits) {
		t.Errorf("err = %v, want ErrInvalidUnits", err)
	}
}

func TestWithHTTPClient(t *testing.T) {
	hc := &http.Client{Timeout: 1 * time.Second}
	c, err := NewClient(WithHTTPClient(hc))
	if err != nil {
		t.Fatalf("error = %v", err)
	}
	if c.HTTPClient != hc {
		t.Error("HTTPClient was not overridden")
	}
	if _, err := NewClient(WithHTTPClient(nil)); err == nil {
		t.Error("WithHTTPClient(nil) should return an error")
	}
}

func TestMustNewClient_PanicsOnInvalid(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic from MustNewClient")
		}
	}()
	_ = MustNewClient(WithUnits("invalid"))
}
