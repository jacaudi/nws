package nwsgo

// RadarStationResponse represents the radar station details.
type RadarStationResponse struct {
	URL      string `json:"@id"`
	Geometry string `json:"geometry"`
	Name     string `json:"name"`
	TimeZone string `json:"timeZone"`
	RDA      RDA    `json:"rda"`
}

type RDA struct {
	Properties RDAProperties `json:"properties"`
}

type RDAProperties struct {
	VolumeCoveragePattern string `json:"volumeCoveragePattern"`
}
