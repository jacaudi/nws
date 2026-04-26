package nws

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestPackageLevelWrappers_DelegateToDefaultClient verifies that the
// back-compat package-level functions (GetPoints, etc.) actually route
// through DefaultClient. We mutate DefaultClient to point at httptest
// and assert the request hits the expected path.
//
// This test mutates the package-level DefaultClient. Do not call
// t.Parallel() here, and do not parallelize any sibling test that
// touches DefaultClient.
func TestPackageLevelWrappers_DelegateToDefaultClient(t *testing.T) {
	hits := map[string]bool{}
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hits[r.URL.Path] = true
		_, _ = w.Write([]byte(`{}`))
	}))
	defer srv.Close()

	saved := DefaultClient
	t.Cleanup(func() { DefaultClient = saved })

	c, err := NewClient(WithBaseURL(srv.URL))
	if err != nil {
		t.Fatal(err)
	}
	DefaultClient = c

	if _, err := GetPoints("47,-122"); err != nil {
		t.Errorf("GetPoints: %v", err)
	}
	if _, err := RadarStations(); err != nil {
		t.Errorf("RadarStations: %v", err)
	}
	if _, err := RadarStation("KATX"); err != nil {
		t.Errorf("RadarStation: %v", err)
	}
	if _, err := GetForecast("SEW", "124,67"); err != nil {
		t.Errorf("GetForecast: %v", err)
	}
	if _, err := GetActiveAlerts(); err != nil {
		t.Errorf("GetActiveAlerts: %v", err)
	}
	if _, err := GetLatestObservations("KBFI"); err != nil {
		t.Errorf("GetLatestObservations: %v", err)
	}

	wantPaths := []string{
		"/points/47,-122",
		"/radar/stations",
		"/radar/stations/KATX",
		"/gridpoints/SEW/124,67/forecast",
		"/alerts/active",
		"/stations/KBFI/observations/latest",
	}
	for _, p := range wantPaths {
		if !hits[p] {
			t.Errorf("expected hit on %s", p)
		}
	}
	if got := len(hits); got != 6 {
		t.Errorf("hits = %d distinct paths, want 6 (%v)", got, hits)
	}
}
