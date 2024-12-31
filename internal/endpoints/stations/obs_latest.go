package stations

import "time"

// https://www.weather.gov/documentation/services-web-api#/default/station_observation_latest

type LatestObservationsResponse struct {
	Context                   Context                   `json:"@context"`
	ID                        string                    `json:"@id"`
	Type                      string                    `json:"@type"`
	Geometry                  string                    `json:"geometry"`
	Elevation                 Elevation                 `json:"elevation"`
	Station                   string                    `json:"station"`
	Timestamp                 time.Time                 `json:"timestamp"`
	RawMessage                string                    `json:"rawMessage"`
	TextDescription           string                    `json:"textDescription"`
	Icon                      string                    `json:"icon"`
	PresentWeather            []interface{}             `json:"presentWeather"`
	Temperature               Temperature               `json:"temperature"`
	Dewpoint                  Dewpoint                  `json:"dewpoint"`
	WindDirection             WindDirection             `json:"windDirection"`
	WindSpeed                 WindSpeed                 `json:"windSpeed"`
	WindGust                  WindGust                  `json:"windGust"`
	BarometricPressure        BarometricPressure        `json:"barometricPressure"`
	SeaLevelPressure          SeaLevelPressure          `json:"seaLevelPressure"`
	Visibility                Visibility                `json:"visibility"`
	MaxTemperatureLast24Hours MaxTemperatureLast24Hours `json:"maxTemperatureLast24Hours"`
	MinTemperatureLast24Hours MinTemperatureLast24Hours `json:"minTemperatureLast24Hours"`
	PrecipitationLastHour     PrecipitationLastHour     `json:"precipitationLastHour"`
	PrecipitationLast3Hours   PrecipitationLast3Hours   `json:"precipitationLast3Hours"`
	PrecipitationLast6Hours   PrecipitationLast6Hours   `json:"precipitationLast6Hours"`
	RelativeHumidity          RelativeHumidity          `json:"relativeHumidity"`
	WindChill                 WindChill                 `json:"windChill"`
	HeatIndex                 HeatIndex                 `json:"heatIndex"`
	CloudLayers               []CloudLayer              `json:"cloudLayers"`
}

type Context struct {
	Version          string           `json:"@version"`
	Wx               string           `json:"wx"`
	S                string           `json:"s"`
	Geo              string           `json:"geo"`
	Unit             string           `json:"unit"`
	Vocab            string           `json:"@vocab"`
	Geometry         Geometry         `json:"geometry"`
	City             string           `json:"city"`
	State            string           `json:"state"`
	Distance         Distance         `json:"distance"`
	Bearing          Bearing          `json:"bearing"`
	Value            Value            `json:"value"`
	UnitCode         UnitCode         `json:"unitCode"`
	ForecastOffice   ForecastOffice   `json:"forecastOffice"`
	ForecastGridData ForecastGridData `json:"forecastGridData"`
	PublicZone       PublicZone       `json:"publicZone"`
	County           County           `json:"county"`
}

type Geometry struct {
	ID   string `json:"@id"`
	Type string `json:"@type"`
}

type Distance struct {
	ID   string `json:"@id"`
	Type string `json:"@type"`
}

type Bearing struct {
	Type string `json:"@type"`
}

type Value struct {
	ID string `json:"@id"`
}

type UnitCode struct {
	ID   string `json:"@id"`
	Type string `json:"@type"`
}

type ForecastOffice struct {
	Type string `json:"@type"`
}

type ForecastGridData struct {
	Type string `json:"@type"`
}

type PublicZone struct {
	Type string `json:"@type"`
}

type County struct {
	Type string `json:"@type"`
}

type Elevation struct {
	UnitCode string  `json:"unitCode"`
	Value    float64 `json:"value"`
}

type Temperature struct {
	UnitCode       string  `json:"unitCode"`
	Value          float64 `json:"value"`
	QualityControl string  `json:"qualityControl"`
}

type Dewpoint struct {
	UnitCode       string  `json:"unitCode"`
	Value          float64 `json:"value"`
	QualityControl string  `json:"qualityControl"`
}

type WindDirection struct {
	UnitCode       string  `json:"unitCode"`
	Value          float64 `json:"value"`
	QualityControl string  `json:"qualityControl"`
}

type WindSpeed struct {
	UnitCode       string  `json:"unitCode"`
	Value          float64 `json:"value"`
	QualityControl string  `json:"qualityControl"`
}

type WindGust struct {
	UnitCode       string   `json:"unitCode"`
	Value          *float64 `json:"value"`
	QualityControl string   `json:"qualityControl"`
}

type BarometricPressure struct {
	UnitCode       string  `json:"unitCode"`
	Value          float64 `json:"value"`
	QualityControl string  `json:"qualityControl"`
}

type SeaLevelPressure struct {
	UnitCode       string  `json:"unitCode"`
	Value          float64 `json:"value"`
	QualityControl string  `json:"qualityControl"`
}

type Visibility struct {
	UnitCode       string  `json:"unitCode"`
	Value          float64 `json:"value"`
	QualityControl string  `json:"qualityControl"`
}

type MaxTemperatureLast24Hours struct {
	UnitCode string   `json:"unitCode"`
	Value    *float64 `json:"value"`
}

type MinTemperatureLast24Hours struct {
	UnitCode string   `json:"unitCode"`
	Value    *float64 `json:"value"`
}

type PrecipitationLastHour struct {
	UnitCode       string   `json:"unitCode"`
	Value          *float64 `json:"value"`
	QualityControl string   `json:"qualityControl"`
}

type PrecipitationLast3Hours struct {
	UnitCode       string   `json:"unitCode"`
	Value          *float64 `json:"value"`
	QualityControl string   `json:"qualityControl"`
}

type PrecipitationLast6Hours struct {
	UnitCode       string   `json:"unitCode"`
	Value          *float64 `json:"value"`
	QualityControl string   `json:"qualityControl"`
}

type RelativeHumidity struct {
	UnitCode       string  `json:"unitCode"`
	Value          float64 `json:"value"`
	QualityControl string  `json:"qualityControl"`
}

type WindChill struct {
	UnitCode       string   `json:"unitCode"`
	Value          *float64 `json:"value"`
	QualityControl string   `json:"qualityControl"`
}

type HeatIndex struct {
	UnitCode       string   `json:"unitCode"`
	Value          *float64 `json:"value"`
	QualityControl string   `json:"qualityControl"`
}

type CloudLayer struct {
	Base   Elevation `json:"base"`
	Amount string    `json:"amount"`
}
