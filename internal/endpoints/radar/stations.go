package radar

// RadarStationsResponse represents the details of all the sites.
// https://www.weather.gov/documentation/services-web-api#/default/radar_stations

type RadarStationsResponse struct {
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
	Type        string      `json:"type"`
	Features    Features    `json:"features"`
}

type Features struct {
	ID                string            `json:"id"`
	Type              string            `json:"type"`
	Geometry          string            `json:"geometry"`
	FeatureProperties FeatureProperties `json:"properties"`
}

type FeatureProperties struct {
	URL         string    `json:"@id"`
	Type        string    `json:"@type"`
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	StationType string    `json:"stationType"`
	Elevation   UnitValue `json:"elevation"`
	TimeZone    string    `json:"timeZone"`
	Latency     Latency   `json:"latency"`
	RDA         RDA       `json:"rda"`
}
