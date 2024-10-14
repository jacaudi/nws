//https://www.weather.gov/documentation/services-web-api
// Package nws implements a basic wrapper around api.weather.gov to grab
// HTTP responses to endpoints (i.e. forecast, alert and radar data) by
// the National Weather Service, an agency of the United States.

package nwsgo

import (
	"encoding/json"
)

// Debug
var debug = false

// GetRadarStation fetches the radar station details for a given station ID.
func RadarStations() (*RadarStationsResponse, error) {
	url := config.endpointRadarStations()
	body, err := config.httpRequest(url)
	if err != nil {
		return nil, err
	}

	var radarStations RadarStationsResponse
	err = json.Unmarshal(body, &radarStations)

	return &radarStations, nil
}

// GetRadarStation fetches the radar station details for a given station ID.
func RadarStation(stationID string) (*RadarStationResponse, error) {
	url := config.endpointRadarStation(stationID)
	body, err := config.httpRequest(url)
	if err != nil {
		return nil, err
	}

	var radarStation RadarStationResponse
	err = json.Unmarshal(body, &radarStation)

	return &radarStation, nil
}
