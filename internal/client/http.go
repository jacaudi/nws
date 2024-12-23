package client

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// Config describes important values for the NWS API and handles making HTTP requests.
type Config struct {
	BaseURL   string `json:"baseUrl"` // Do not include a trailing slash
	UserAgent string `json:"userAgent"`
	Accept    string `json:"accept"`
	Units     string `json:"units"`
}

// GetDefaultConfig returns the default configuration for the weather.gov API.
func GetDefaultConfig() Config {
	return Config{
		BaseURL:   "https://api.weather.gov",
		UserAgent: "nwsgo/0.0.2 (+https://github.com/jacaudi/nwsgo)",
		Accept:    "application/ld+json",
		Units:     "", // Defaults to US units if unspecified
	}
}

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
