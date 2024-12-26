package gridpoints

import (
	"time"
)

// https://www.weather.gov/documentation/services-web-api#/default/gridpoint_forecast

type ForecastResponse struct {
	Context           Context   `json:"@context"`
	Geometry          string    `json:"geometry"`
	Units             string    `json:"units"`
	ForecastGenerator string    `json:"forecastGenerator"`
	GeneratedAt       time.Time `json:"generatedAt"`
	UpdateTime        time.Time `json:"updateTime"`
	ValidTimes        string    `json:"validTimes"`
	Elevation         Elevation `json:"elevation"`
	Periods           []Period  `json:"periods"`
}

type Context struct {
	Version string `json:"@version"`
	Wx      string `json:"wx"`
	Geo     string `json:"geo"`
	Unit    string `json:"unit"`
	Vocab   string `json:"@vocab"`
}

type Elevation struct {
	UnitCode string  `json:"unitCode"`
	Value    float64 `json:"value"`
}

type Period struct {
	Number                     int           `json:"number"`
	Name                       string        `json:"name"`
	StartTime                  time.Time     `json:"startTime"`
	EndTime                    time.Time     `json:"endTime"`
	IsDaytime                  bool          `json:"isDaytime"`
	Temperature                int           `json:"temperature"`
	TemperatureUnit            string        `json:"temperatureUnit"`
	TemperatureTrend           string        `json:"temperatureTrend"`
	ProbabilityOfPrecipitation Precipitation `json:"probabilityOfPrecipitation"`
	WindSpeed                  string        `json:"windSpeed"`
	WindDirection              string        `json:"windDirection"`
	Icon                       string        `json:"icon"`
	ShortForecast              string        `json:"shortForecast"`
	DetailedForecast           string        `json:"detailedForecast"`
}

type Precipitation struct {
	UnitCode string `json:"unitCode"`
	Value    int    `json:"value"`
}
