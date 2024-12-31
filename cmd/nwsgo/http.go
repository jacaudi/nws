package nwsgo

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

// httpRequest makes an HTTP request to the NWS API and returns the response body.
func (c *Config) httpRequest(url string, httpResponse bool) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("User-Agent", c.UserAgent)
	req.Header.Set("Accept", c.Accept)

	if debug {
		log.Printf("Making request to URL: %s", url)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if httpResponse {
		log.Printf("Received response status: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if httpResponse {
		log.Printf("Response body: %s", string(body))
	}

	return body, nil
}

// SetUserAgent changes the User-Agent header for API requests.
func (c *Config) SetUserAgent(userAgent string) {
	if len(userAgent) == 0 {
		panic("The API requires a User-Agent")
	}
	c.UserAgent = userAgent
}

// SetUnits changes the unit system (US or SI) for API requests.
func (c *Config) SetUnits(units string) {
	units = strings.ToLower(units)
	if units != "us" && units != "si" {
		c.Units = ""
	} else {
		c.Units = units
	}
}

// Endpoints
func (c *Config) endpointPoints(latlon string) string {
	return fmt.Sprintf("%s/points/%s", c.BaseURL, latlon)
}

func (c *Config) endpointRadarStations() string {
	return fmt.Sprintf("%s/radar/stations", c.BaseURL)
}

func (c *Config) endpointRadarStation(stationID string) string {
	return fmt.Sprintf("%s/radar/stations/%s", c.BaseURL, stationID)
}

func (c *Config) endpointGridForecast(wfo string, gridpoint string) string {
	return fmt.Sprintf("%s/gridpoints/%s/%s/forecast", c.BaseURL, wfo, gridpoint)
}

func (c *Config) endpointActiveAlerts() string {
	return fmt.Sprintf("%s/alerts/active", c.BaseURL)
}
