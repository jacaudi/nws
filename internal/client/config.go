package client

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
