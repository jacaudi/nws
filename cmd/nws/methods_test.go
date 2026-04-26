package nws

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClient_endpointBuilders(t *testing.T) {
	c, err := NewClient(WithBaseURL("https://example.com"))
	if err != nil {
		t.Fatal(err)
	}
	cases := []struct {
		name string
		got  string
		want string
	}{
		{"points", c.endpointPoints("47,-122"), "/points/47,-122"},
		{"radarStations", c.endpointRadarStations(), "/radar/stations"},
		{"radarStation", c.endpointRadarStation("KATX"), "/radar/stations/KATX"},
		{"forecast", c.endpointGridForecast("SEW", "124,67"), "/gridpoints/SEW/124,67/forecast"},
		{"alerts", c.endpointActiveAlerts(), "/alerts/active"},
		{"obs", c.endpointLatestObservations("KBFI"), "/stations/KBFI/observations/latest"},
	}
	for _, tc := range cases {
		if tc.got != tc.want {
			t.Errorf("%s: got %q, want %q", tc.name, tc.got, tc.want)
		}
	}
}

func TestClient_RadarStation_HappyPath(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/radar/stations/KATX" {
			t.Errorf("path = %q, want /radar/stations/KATX", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/ld+json")
		_, _ = w.Write([]byte(`{"name":"Seattle"}`))
	}))
	defer srv.Close()

	c, err := NewClient(WithBaseURL(srv.URL))
	if err != nil {
		t.Fatal(err)
	}

	resp, err := c.RadarStation(context.Background(), "KATX")
	if err != nil {
		t.Fatalf("RadarStation: %v", err)
	}
	if resp == nil {
		t.Fatal("response is nil")
	}
	if resp.Name != "Seattle" {
		t.Errorf("Name = %q, want Seattle", resp.Name)
	}
}

func TestClient_GetActiveAlerts_HappyPath(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/alerts/active" {
			t.Errorf("path = %q", r.URL.Path)
		}
		_, _ = w.Write([]byte(`{"@graph":[]}`))
	}))
	defer srv.Close()

	c, err := NewClient(WithBaseURL(srv.URL))
	if err != nil {
		t.Fatal(err)
	}

	if _, err := c.GetActiveAlerts(context.Background()); err != nil {
		t.Errorf("GetActiveAlerts: %v", err)
	}
}
