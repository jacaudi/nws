package alerts

// https://www.weather.gov/documentation/services-web-api#/default/alerts_active

type ActiveAlertsResponse struct {
	Context []interface{} `json:"@context"`
	Type    string        `json:"type"`
}
