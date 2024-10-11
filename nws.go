//https://www.weather.gov/documentation/services-web-api#/default/radar_stations

// Package nws implements a basic wrapper around api.weather.gov to
// grab HTTP responses to endpoints (i.e.: weather & forecast data)
// by the National Weather Service, an agency of the United States.
package nwsgo

import (
	"encoding/json"
	"fmt"
	"log"
)

// GetRadarStation fetches the radar station details for a given station ID.
func RadarStation(stationID string) (*RadarStationResponse, error) {
	url := config.endpointRadarStation(stationID)
	body, err := config.MakeRequest(url)
	if err != nil {
		return nil, err
	}

	if debug {
		var prettyJSON map[string]interface{}
		if err := json.Unmarshal(body, &prettyJSON); err == nil {
			prettyBody, _ := json.MarshalIndent(prettyJSON, "", "  ")
			log.Printf("\n-----------------------------\nPretty-printed response body:\n%s\n-----------------------------\n", string(prettyBody))
		} else {
			log.Printf("\n-----------------------------\nFailed to pretty-print response body:\n%s\n-----------------------------\n", err)
		}
	}

	var radarStation RadarStationResponse
	err = json.Unmarshal(body, &radarStation)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if debug {
		log.Printf("\n-----------------------------\nParsed RadarStation:\n%+v\n-----------------------------\n", radarStation)
	}

	return &radarStation, nil
}
