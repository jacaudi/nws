package radar

// RadarStationResponse represents the site details of a single radar site
// https://www.weather.gov/documentation/services-web-api#/default/radar_station

type RadarStationResponse struct {
	Adaptation  Adaptation  `json:"adaptation"`
	Context     Context     `json:"@context"`
	Elevation   UnitValue   `json:"elevation"`
	Geometry    string      `json:"geometry"`
	Latency     Latency     `json:"latency"`
	Name        string      `json:"name"`
	Performance Performance `json:"performance"`
	RDA         RDA         `json:"rda"`
	StationType string      `json:"stationType"`
	TimeZone    string      `json:"timeZone"`
	URL         string      `json:"@id"`
}
