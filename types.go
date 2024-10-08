import (
    "encoding/json"
    "fmt"
)

// ContextType can handle both strings and objects in the @context field
type ContextType struct {
    Contexts []interface{}
}

// UnmarshalJSON custom unmarshaller for ContextType
func (c *ContextType) UnmarshalJSON(data []byte) error {
    var singleContext interface{}
    if err := json.Unmarshal(data, &singleContext); err == nil {
        c.Contexts = append(c.Contexts, singleContext)
        return nil
    }

    var multipleContexts []interface{}
    if err := json.Unmarshal(data, &multipleContexts); err == nil {
        c.Contexts = multipleContexts
        return nil
    }

    return fmt.Errorf("failed to unmarshal @context field")
}

// RadarStationListResponse holds the JSON values from /radar/stations
type RadarStationListResponse struct {
    Context  ContextType               `json:"@context"`
    Type     string                    `json:"type"`
    Features []RadarStationListFeature `json:"features"`
}

// RadarStationListFeature represents an individual radar station feature
type RadarStationListFeature struct {
    ID         string                     `json:"id"`
    Type       string                     `json:"type"`
    Geometry   RadarStationGeometry       `json:"geometry"`
    Properties RadarStationListProperties `json:"properties"`
}

// RadarStationGeometry represents the geometry of a radar station
type RadarStationGeometry struct {
    Type        string    `json:"type"`
    Coordinates []float64 `json:"coordinates"`
}

// RadarStationListProperties holds detailed properties of a radar station
type RadarStationListProperties struct {
    ID             string          `json:"@id"`
    Type           string          `json:"@type"`
    StationID      string          `json:"id"`
    StationName    string          `json:"name"`
    StationType    string          `json:"stationType"`
    Elevation      Elevation       `json:"elevation"`
    TimeZone       string          `json:"timeZone"`
    Latency        Latency         `json:"latency"`
    RDA            RDA             `json:"rda"`
}

// Elevation represents the elevation details of a radar station
type Elevation struct {
    UnitCode string  `json:"unitCode"`
    Value    float64 `json:"value"`
}

// Latency represents the latency details of a radar station
type Latency struct {
    Current                Measurement `json:"current"`
    Average                Measurement `json:"average"`
    Max                    Measurement `json:"max"`
    LevelTwoLastReceivedTime string    `json:"levelTwoLastReceivedTime"`
    MaxLatencyTime         string      `json:"maxLatencyTime"`
    ReportingHost          string      `json:"reportingHost"`
    Host                   string      `json:"host"`
}

// Measurement represents a measurement with a unit code and value
type Measurement struct {
    UnitCode string  `json:"unitCode"`
    Value    float64 `json:"value"`
}

// RDA represents the RDA details of a radar station
type RDA struct {
    Timestamp      string         `json:"timestamp"`
    ReportingHost  string         `json:"reportingHost"`
    Properties     RDAProperties  `json:"properties"`
}

// RDAProperties represents the properties of the RDA
type RDAProperties struct {
    ResolutionVersion              *string         `json:"resolutionVersion"`
    Nl2Path                        string          `json:"nl2Path"`
    VolumeCoveragePattern          string          `json:"volumeCoveragePattern"`
    ControlStatus                  string          `json:"controlStatus"`
    BuildNumber                    float64         `json:"buildNumber"`
    AlarmSummary                   string          `json:"alarmSummary"`
    Mode                           string          `json:"mode"`
    GeneratorState                 string          `json:"generatorState"`
    SuperResolutionStatus          string          `json:"superResolutionStatus"`
    OperabilityStatus              string          `json:"operabilityStatus"`
    Status                         string          `json:"status"`
    AverageTransmitterPower        Measurement     `json:"averageTransmitterPower"`
    ReflectivityCalibrationCorrection Measurement `json:"reflectivityCalibrationCorrection"`
}
