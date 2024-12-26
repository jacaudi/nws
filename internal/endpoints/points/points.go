package points

// https://www.weather.gov/documentation/services-web-api#/default/point

// PointsResponse is a custom struct type for storing forecast data for a specific location.
type PointsResponse struct {
	Context             Context          `json:"@context"`
	ID                  string           `json:"@id"`
	Type                string           `json:"@type"`
	Geometry            string           `json:"geometry"`
	CWA                 string           `json:"cwa"`
	ForecastOffice      string           `json:"forecastOffice"`
	GridID              string           `json:"gridId"`
	GridX               float64          `json:"gridX"`
	GridY               float64          `json:"gridY"`
	ForecastURL         string           `json:"forecast"`
	ForecastHourlyURL   string           `json:"forecastHourly"`
	ForecastGridData    string           `json:"forecastGridData"`
	ObservationStations string           `json:"observationStations"`
	RelativeLocation    RelativeLocation `json:"relativeLocation"`
	ForecastZone        string           `json:"forecastZone"`
	County              string           `json:"county"`
	FireWeatherZone     string           `json:"fireWeatherZone"`
	TimeZone            string           `json:"timeZone"`
	RadarStation        string           `json:"radarStation"`
}

type Context struct {
	Version string `json:"@version"`
	Wx      string `json:"wx"`
	S       string `json:"s"`
	Geo     string `json:"geo"`
	Unit    string `json:"unit"`
	Vocab   string `json:"@vocab"`
}

// RelativeLocation is a custom struct type for storing relative location information, such as the city and state.
type RelativeLocation struct {
	City     string   `json:"city"`
	State    string   `json:"state"`
	Geometry string   `json:"geometry"`
	Distance Distance `json:"distance"`
	Bearing  Bearing  `json:"bearing"`
}

// Distance is a custom struct type for storing distance information.
type Distance struct {
	UnitCode string  `json:"unitCode"`
	Value    float64 `json:"value"`
}

// Bearing is a custom struct type for storing bearing information.
type Bearing struct {
	UnitCode string  `json:"unitCode"`
	Value    float64 `json:"value"`
}
