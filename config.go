package nws-go

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

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
		UserAgent: "github.com/jacaudi/nws",
		Accept:    "application/ld+json",
		Units:     "", // Defaults to US units if unspecified
	}
}

// Make a GET request and decode the response into the provided reference.
func (c *Config) GetAndDecode(endpoint string, v any) error {
	// Create and configure the HTTP request
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return err
	}

	req.Header.Add("Accept", c.Accept)
	req.Header.Add("User-Agent", c.UserAgent)
	req.Header.Add("feature-flags", "forecast_temperature_qv, forecast_wind_speed_qv")

	// Send the HTTP request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	// Check if the status code is OK
	if res.StatusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("%d %s", res.StatusCode, res.Status))
	}

	// Decode the JSON response into the provided struct
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(v); err != nil {
		return err
	}

	return nil
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
	return fmt.Sprintf("%s/radar/stations/%s", c.BaseURL, stationID)
}

func (c *Config) endpointRadarStationAlarms(stationID string) string {
	return fmt.Sprintf("%s/radar/stations/%s/alarms", c.BaseURL, stationID)
}
