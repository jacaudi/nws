//https://www.weather.gov/documentation/services-web-api

// Package nws is a Go client for the U.S. National Weather Service API
// (api.weather.gov). It exposes a *Client type for explicit
// configuration plus package-level wrapper functions that delegate to
// DefaultClient for one-shot calls.
//
// Idiomatic usage:
//
//	client, err := nws.NewClient(
//	    nws.WithUserAgent("myapp/1.0 (contact@example.com)"),
//	)
//	if err != nil { return err }
//	station, err := client.RadarStation(ctx, "KATX")
//
// One-shot usage (uses DefaultClient with context.Background()):
//
//	station, err := nws.RadarStation("KATX")
package nws

import (
	"context"

	"github.com/jacaudi/nws/internal/endpoints/alerts"
	"github.com/jacaudi/nws/internal/endpoints/gridpoints"
	"github.com/jacaudi/nws/internal/endpoints/points"
	"github.com/jacaudi/nws/internal/endpoints/radar"
	"github.com/jacaudi/nws/internal/endpoints/stations"
)

// GetPoints is the DefaultClient.GetPoints wrapper.
// See (*Client).GetPoints.
func GetPoints(latlon string) (*points.PointsResponse, error) {
	return DefaultClient.GetPoints(context.Background(), latlon)
}

// RadarStations is the DefaultClient.RadarStations wrapper.
// See (*Client).RadarStations.
func RadarStations() (*radar.RadarStationsResponse, error) {
	return DefaultClient.RadarStations(context.Background())
}

// RadarStation is the DefaultClient.RadarStation wrapper.
// See (*Client).RadarStation.
func RadarStation(stationID string) (*radar.RadarStationResponse, error) {
	return DefaultClient.RadarStation(context.Background(), stationID)
}

// GetForecast is the DefaultClient.GetForecast wrapper.
// See (*Client).GetForecast.
func GetForecast(wfo, gridpoint string) (*gridpoints.ForecastResponse, error) {
	return DefaultClient.GetForecast(context.Background(), wfo, gridpoint)
}

// GetActiveAlerts is the DefaultClient.GetActiveAlerts wrapper.
// See (*Client).GetActiveAlerts.
func GetActiveAlerts() (*alerts.ActiveAlertsResponse, error) {
	return DefaultClient.GetActiveAlerts(context.Background())
}

// GetLatestObservations is the DefaultClient.GetLatestObservations
// wrapper. See (*Client).GetLatestObservations.
func GetLatestObservations(stationID string) (*stations.LatestObservationsResponse, error) {
	return DefaultClient.GetLatestObservations(context.Background(), stationID)
}
