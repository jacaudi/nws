package points

// https://www.weather.gov/documentation/services-web-api#/default/point

// ForecastData is a custom struct type for storing forecast data for a specific location.
type PointsResponse struct {
	ID                  string  `json:"@id"`
	ForecastOffice      string  `json:"forecastOffice"`
	GridID              string  `json:"gridId"`
	GridX               float64 `json:"gridX"`
	GridY               float64 `json:"gridY"`
	ForecastURL         string  `json:"forecast"`
	ForecastHourlyURL   string  `json:"forecastHourly"`
	ForecastGridData    string  `json:"forecastGridData"`
	ObservationStations string  `json:"observationStations"`
	ForecastZone        string  `json:"forecastZone"`
	County              string  `json:"county"`
	FireWeatherZone     string  `json:"fireWeatherZone"`
	TimeZone            string  `json:"timeZone"`
	RadarStation        string  `json:"radarStation"`
}

// Point is a custom struct type for storing information about a location, such as latitude and longitude.
type Point struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lon"`
}

// RelativeLocation is a custom struct type for storing relative location information, such as the city and state.
type RelativeLocation struct {
	City     string  `json:"city"`
	State    string  `json:"state"`
	Distance float64 `json:"distance"`
	Bearing  float64 `json:"bearing"`
}
