package nwsgo

// RadarStationResponse represents the radar station details.
type RadarStationResponse struct {
	URL      string `json:"@id"`
	Geometry string `json:"geometry"` // Assuming geometry is a string
	Name     string `json:"name"`
	TimeZone string `json:"timeZone"`
	RDA      struct {
		Properties struct {
			VolumeCoveragePattern string `json:"volumeCoveragePattern"` // Changed to int
		} `json:"properties"`
	} `json:"rda"`
}
