package alerts

import "time"

// https://www.weather.gov/documentation/services-web-api#/default/alerts_active

type ActiveAlertsResponse struct {
	Type    string    `json:"type"`
	Title   string    `json:"title"`
	Updated time.Time `json:"updated"`
	Data    []Data    `json:"@graph"`
}

type Data struct {
	ID              string      `json:"@id"`
	Type            string      `json:"type"`
	Geometry        *string     `json:"geometry"`
	AreaDescription string      `json:"areaDesc"`
	Sent            time.Time   `json:"sent"`
	Effective       time.Time   `json:"effective"`
	Onset           time.Time   `json:"onset"`
	Expires         time.Time   `json:"expires"`
	Ends            time.Time   `json:"ends"`
	Status          string      `json:"status"`
	MessageType     string      `json:"messageType"`
	Category        string      `json:"category"`
	Severity        string      `json:"severity"`
	Certainty       string      `json:"certainty"`
	Urgency         string      `json:"urgency"`
	Event           string      `json:"event"`
	Sender          string      `json:"sender"`
	SenderName      string      `json:"senderName"`
	Headline        string      `json:"headline"`
	Description     string      `json:"description"`
	Instruction     string      `json:"instruction"`
	Response        string      `json:"response"`
	Parameters      Parameters  `json:"parameters"`
	Geocode         Geocode     `json:"geocode"`
	AffectedZones   []string    `json:"affectedZones"`
	References      []Reference `json:"references"`
}

type Geocode struct {
	SAME []string `json:"SAME"`
	UGC  []string `json:"UGC"`
}

type Reference struct {
	ID         string `json:"@id"`
	Identifier string `json:"identifier"`
	Sender     string `json:"sender"`
	Sent       string `json:"sent"`
}

type Parameters struct {
	AWIPSidentifier   []string    `json:"AWIPSidentifier"`
	WMOidentifier     []string    `json:"WMOidentifier"`
	NWSheadline       []string    `json:"NWSheadline"`
	BLOCKCHANNEL      []string    `json:"BLOCKCHANNEL"`
	VTEC              []string    `json:"VTEC"`
	EventEndingTime   []time.Time `json:"eventEndingTime"`
	ExpiredReferences []string    `json:"expiredReferences"`
}
