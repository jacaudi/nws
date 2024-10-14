package nwsgo

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

// Show the HTTP response
var httpResponse = false

// Config instance for the API calls executed by the NWS client.
var config = GetDefaultConfig()

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
		UserAgent: "github.com/jacaudi/nwsgo",
		Accept:    "application/ld+json",
		Units:     "", // Defaults to US units if unspecified
	}
}

// httpRequest makes an HTTP request to the NWS API and returns the response body.
func (c *Config) httpRequest(url string) ([]byte, error) {
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
func (c *Config) endpointRadarStations() string {
	return fmt.Sprintf("%s/radar/stations", c.BaseURL)
}

func (c *Config) endpointRadarStation(stationID string) string {
	return fmt.Sprintf("%s/radar/station/%s", c.BaseURL, stationID)
}
