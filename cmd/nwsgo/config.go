package nws

// Config instance for the API calls executed by the NWS client.
var config = getDefaultConfig(Config{})

// Config describes important values for the NWS API and handles making HTTP requests.
type Config struct {
	BaseURL   string `json:"baseUrl"` // Do not include a trailing slash
	UserAgent string `json:"userAgent"`
	Accept    string `json:"accept"`
	Units     string `json:"units"`
}

// GetDefaultConfig returns the default configuration for the weather.gov API,
// optionally merging it with a custom configuration.
func getDefaultConfig(custom Config) Config {
	defaultConfig := Config{
		BaseURL:   "https://api.weather.gov",
		UserAgent: "nws/0.0.2 (+https://github.com/jacaudi/nws)",
		Accept:    "application/ld+json",
		Units:     "", // Defaults to US units if unspecified
	}

	if custom.UserAgent != "" {
		defaultConfig.UserAgent = custom.UserAgent
	}
	if custom.Accept != "" {
		defaultConfig.Accept = custom.Accept
	}
	if custom.Units != "" {
		defaultConfig.Units = custom.Units
	}

	return defaultConfig
}
