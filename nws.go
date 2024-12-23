//https://www.weather.gov/documentation/services-web-api
// Package nws implements a basic wrapper around api.weather.gov to grab
// HTTP responses to endpoints (i.e. forecast, alert and radar data) by
// the National Weather Service, an agency of the United States.

package nwsgo

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/jacaudi/nwsgo/internal/client"
	"github.com/jacaudi/nwsgo/internal/endpoints/points"
	"github.com/jacaudi/nwsgo/internal/endpoints/radar"
)

// Debug
var debug = true
var c = client.GetDefaultConfig()

// httpRequest makes an HTTP request to the NWS API and returns the response body.
func httpRequest(url string, agent string, accept string, units string, debug bool) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Accept", accept)

	if len(agent) == 0 {
		log.Panicf("The NWS API requires a User-Agent")
	} else {
		req.Header.Set("User-Agent", agent)
	}

	if len(units) == 0 {
		req.Header.Set("Units", "") // Defaults to US units if unspecified
	} else {
		req.Header.Set("Units", units)
	}

	if debug {
		log.Printf("Making request to URL: %s", url)
		log.Printf("Header: %s", accept)
		log.Printf("Header: %s", agent)
		log.Printf("Header: %s", units)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if debug {
		log.Printf("Received response status: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if debug {
		log.Printf("Response body: %s", string(body))
	}

	return body, nil
}

// GetRadarStation fetches the radar station details for a given station ID.
func GetPoints(latlon string) (*points.PointsResponse, error) {
	url := fmt.Sprintf("%s/points/%s", c.BaseURL, latlon)
	response, err := httpRequest(url, c.UserAgent, c.Accept, c.Units, debug)
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
	response, err := httpRequest(url, c.UserAgent, c.Accept, c.Units, debug)
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
	response, err := httpRequest(url, c.UserAgent, c.Accept, c.Units, debug)
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
