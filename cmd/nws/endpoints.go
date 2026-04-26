package nws

import "fmt"

// endpoint URL builders. Kept private; consumers should call the methods
// on *Client (GetPoints, RadarStation, etc.) which call these internally.

func (c *Client) endpointPoints(latlon string) string {
	return fmt.Sprintf("/points/%s", latlon)
}

func (c *Client) endpointRadarStations() string {
	return "/radar/stations"
}

func (c *Client) endpointRadarStation(stationID string) string {
	return fmt.Sprintf("/radar/stations/%s", stationID)
}

func (c *Client) endpointGridForecast(wfo, gridpoint string) string {
	return fmt.Sprintf("/gridpoints/%s/%s/forecast", wfo, gridpoint)
}

func (c *Client) endpointActiveAlerts() string {
	return "/alerts/active"
}

func (c *Client) endpointLatestObservations(stationID string) string {
	return fmt.Sprintf("/stations/%s/observations/latest", stationID)
}
