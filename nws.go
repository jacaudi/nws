//https://www.weather.gov/documentation/services-web-api
// Package nws implements a basic wrapper around api.weather.gov to grab
// HTTP responses to endpoints (i.e. forecast, alert and radar data) by
// the National Weather Service, an agency of the United States.

package nwsgo

import (
	"encoding/json"
	"fmt"

	"github.com/jacaudi/nwsgo/internal/endpoints/gridpoints"
	"github.com/jacaudi/nwsgo/internal/endpoints/points"
	"github.com/jacaudi/nwsgo/internal/endpoints/radar"
)

// Debug
var debug = true

// GetPoints grabs NWS data at the following Lat/Lon coordinates
func GetPoints(latlon string) (*points.PointsResponse, error) {
	url := config.endpointPoints(latlon)
	response, err := config.httpRequest(url)
	if err != nil {
		return nil, fmt.Errorf("failed to make HTTP request: %w", err)
	}
	var pointsResponse points.PointsResponse
	err = json.Unmarshal(response, &pointsResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	return &pointsResponse, nil
}

// GetRadarStation fetches the radar station details for a given station ID.
func RadarStations() (*radar.RadarStationsResponse, error) {
	url := config.endpointRadarStations()
	response, err := config.httpRequest(url)
	if err != nil {
		return nil, fmt.Errorf("failed to make HTTP request: %w", err)
	}
	var radarStationsResponse radar.RadarStationsResponse
	err = json.Unmarshal(response, &radarStationsResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	return &radarStationsResponse, nil
}

// RadarStation fetches the radar station details for a given station ID.
func RadarStation(stationID string) (*radar.RadarStationResponse, error) {
	url := config.endpointRadarStation(stationID)
	response, err := config.httpRequest(url)
	if err != nil {
		return nil, fmt.Errorf("failed to make HTTP request: %w", err)
	}
	var radarStationResponse radar.RadarStationResponse
	err = json.Unmarshal(response, &radarStationResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	return &radarStationResponse, nil
}

// GetForecast fetches the forecast details for a given Lat/Lon.
func GetForecast(wfo string, gridpoint string) (*gridpoints.ForecastResponse, error) {
	url := config.endpointGridForecast(wfo, gridpoint)
	response, err := config.httpRequest(url)
	if err != nil {
		return nil, fmt.Errorf("failed to make HTTP request: %w", err)
	}
	var ForecastResponse gridpoints.ForecastResponse
	err = json.Unmarshal(response, &ForecastResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}
	return &ForecastResponse, nil
}
