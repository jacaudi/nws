package nws

import (
	"context"

	"github.com/jacaudi/nws/internal/endpoints/alerts"
	"github.com/jacaudi/nws/internal/endpoints/gridpoints"
	"github.com/jacaudi/nws/internal/endpoints/points"
	"github.com/jacaudi/nws/internal/endpoints/radar"
	"github.com/jacaudi/nws/internal/endpoints/stations"
)

// GetPoints fetches geographic metadata (grid ID, relative location,
// forecast URL) for a "lat,lon" coordinate string from /points/{latlon}.
func (c *Client) GetPoints(ctx context.Context, latlon string) (*points.PointsResponse, error) {
	var out points.PointsResponse
	if err := c.get(ctx, c.endpointPoints(latlon), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// RadarStations fetches the list of all NWS radar stations from
// /radar/stations.
func (c *Client) RadarStations(ctx context.Context) (*radar.RadarStationsResponse, error) {
	var out radar.RadarStationsResponse
	if err := c.get(ctx, c.endpointRadarStations(), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// RadarStation fetches metadata for a single radar station by its ID
// (e.g. "KATX") from /radar/stations/{stationID}.
func (c *Client) RadarStation(ctx context.Context, stationID string) (*radar.RadarStationResponse, error) {
	var out radar.RadarStationResponse
	if err := c.get(ctx, c.endpointRadarStation(stationID), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetForecast fetches a multi-period text forecast for a WFO + gridpoint
// pair from /gridpoints/{wfo}/{gridpoint}/forecast. Look up wfo and
// gridpoint via GetPoints first.
func (c *Client) GetForecast(ctx context.Context, wfo, gridpoint string) (*gridpoints.ForecastResponse, error) {
	var out gridpoints.ForecastResponse
	if err := c.get(ctx, c.endpointGridForecast(wfo, gridpoint), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetActiveAlerts fetches all currently active NWS alerts (warnings,
// watches, advisories) from /alerts/active.
func (c *Client) GetActiveAlerts(ctx context.Context) (*alerts.ActiveAlertsResponse, error) {
	var out alerts.ActiveAlertsResponse
	if err := c.get(ctx, c.endpointActiveAlerts(), &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// GetLatestObservations fetches the most recent observation for an
// observation station ID (e.g. "KBFI") from
// /stations/{stationID}/observations/latest.
func (c *Client) GetLatestObservations(ctx context.Context, stationID string) (*stations.LatestObservationsResponse, error) {
	var out stations.LatestObservationsResponse
	if err := c.get(ctx, c.endpointLatestObservations(stationID), &out); err != nil {
		return nil, err
	}
	return &out, nil
}
