package nws

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestClient_get_HappyPath(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if got := r.Header.Get("User-Agent"); !strings.HasPrefix(got, "nws-go/") {
			t.Errorf("User-Agent = %q, want prefix nws-go/", got)
		}
		if got := r.Header.Get("Accept"); got != "application/ld+json" {
			t.Errorf("Accept = %q, want application/ld+json", got)
		}
		if r.URL.Path != "/test" {
			t.Errorf("path = %q, want /test", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/ld+json")
		_, _ = w.Write([]byte(`{"name":"hello"}`))
	}))
	defer srv.Close()

	c, err := NewClient(WithBaseURL(srv.URL))
	if err != nil {
		t.Fatal(err)
	}

	var out struct {
		Name string `json:"name"`
	}
	if err := c.get(context.Background(), "/test", &out); err != nil {
		t.Fatalf("get: %v", err)
	}
	if out.Name != "hello" {
		t.Errorf("Name = %q, want hello", out.Name)
	}
}

func TestClient_get_NonOKStatus_ReturnsAPIError(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		_, _ = w.Write([]byte(`{"detail":"not found"}`))
	}))
	defer srv.Close()

	c, err := NewClient(WithBaseURL(srv.URL))
	if err != nil {
		t.Fatal(err)
	}

	var out struct{}
	gotErr := c.get(context.Background(), "/missing", &out)
	var apiErr *APIError
	if !errors.As(gotErr, &apiErr) {
		t.Fatalf("got %T (%v), want *APIError", gotErr, gotErr)
	}
	if apiErr.StatusCode != http.StatusNotFound {
		t.Errorf("StatusCode = %d, want 404", apiErr.StatusCode)
	}
	if apiErr.URL != srv.URL+"/missing" {
		t.Errorf("URL = %q, want %s/missing", apiErr.URL, srv.URL)
	}
	if !strings.Contains(apiErr.Body, "not found") {
		t.Errorf("Body = %q, expected to contain 'not found'", apiErr.Body)
	}
}

func TestClient_get_EmptyUserAgent_ReturnsErr(t *testing.T) {
	c, err := NewClient()
	if err != nil {
		t.Fatal(err)
	}
	c.UserAgent = "" // post-construction mutation simulating the bug case

	var out struct{}
	gotErr := c.get(context.Background(), "/whatever", &out)
	if !errors.Is(gotErr, ErrUserAgentRequired) {
		t.Errorf("err = %v, want ErrUserAgentRequired", gotErr)
	}
}

func TestClient_get_ContextCanceled(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		<-r.Context().Done() // server blocks until the client's ctx cancels
	}))
	defer srv.Close()

	c, err := NewClient(WithBaseURL(srv.URL))
	if err != nil {
		t.Fatal(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	cancel() // cancel before the call

	var out struct{}
	gotErr := c.get(ctx, "/", &out)
	if gotErr == nil {
		t.Fatal("expected error from canceled context, got nil")
	}
	if !errors.Is(gotErr, context.Canceled) {
		t.Errorf("err = %v, want errors.Is(err, context.Canceled)", gotErr)
	}
}

func TestClient_get_ErrorBodyTruncatedTo1KiB(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		// 4 KiB of 'x'
		_, _ = w.Write(bytes.Repeat([]byte("x"), 4096))
	}))
	defer srv.Close()

	c, err := NewClient(WithBaseURL(srv.URL))
	if err != nil {
		t.Fatal(err)
	}

	var out struct{}
	gotErr := c.get(context.Background(), "/big-error", &out)
	var apiErr *APIError
	if !errors.As(gotErr, &apiErr) {
		t.Fatalf("got %T, want *APIError", gotErr)
	}
	if got := len(apiErr.Body); got != 1024 {
		t.Errorf("len(Body) = %d, want 1024", got)
	}
}
