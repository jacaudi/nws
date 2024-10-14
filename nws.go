//https://www.weather.gov/documentation/services-web-api#/default/radar_stations

// Package nws implements a basic wrapper around api.weather.gov to
// grab HTTP responses to endpoints (i.e.: weather & forecast data)
// by the National Weather Service, an agency of the United States.
package nwsgo

import (
	"encoding/json"
)

// Debug
var debug = false

// GetRadarStation fetches the radar station details for a given station ID.
func endpointRadarStations() (*RadarStationsResponse, error) {
	url := config.endpointRadarStation()
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
