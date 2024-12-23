//https://www.weather.gov/documentation/services-web-api
// Package nws implements a basic wrapper around api.weather.gov to grab
// HTTP responses to endpoints (i.e. forecast, alert and radar data) by
// the National Weather Service, an agency of the United States.

package nwsgo

import (
	"encoding/json"
	"fmt"

	"github.com/jacaudi/nwsgo/internal/client"
	"github.com/jacaudi/nwsgo/internal/endpoints/points"
	"github.com/jacaudi/nwsgo/internal/endpoints/radar"
)

// Debug
var debug = true
var c = client.GetDefaultConfig()

// GetRadarStation fetches the radar station details for a given station ID.
func GetPoints(latlon string) (*points.PointsResponse, error) {
	url := fmt.Sprintf("%s/points/%s", c.BaseURL, latlon)
	response, err := client.HttpRequest(url, c.UserAgent, c.Accept, c.Units, debug)
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
	url := fmt.Sprintf("%s/radar/stations", c.BaseURL)
	response, err := client.HttpRequest(url, c.UserAgent, c.Accept, c.Units, debug)
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
	url := fmt.Sprintf("%s/radar/stations/%s", c.BaseURL, stationID)
	response, err := client.HttpRequest(url, c.UserAgent, c.Accept, c.Units, debug)
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
