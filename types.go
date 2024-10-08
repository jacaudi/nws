package nwsgo

import "encoding/json"

// api.weather.gov/radar/stations
// RadarStationListResponse holds the JSON values from /radar/stations
type RadarStationListResponse struct {
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

// RadarStationListProperties holds detailed properties of a radar station
type RadarStationListProperties struct {
	ID          string    `json:"@id"`
	Type        string    `json:"@type"`
	StationID   string    `json:"id"`
	StationName string    `json:"name"`
	StationType string    `json:"stationType"`
	Elevation   Elevation `json:"elevation"`
	TimeZone    string    `json:"timeZone"`
	Latency     Latency   `json:"latency"`
	RDA         RDA       `json:"rda"`
}

// api.weather.gov/radar/stations/{stationId}
// RadarStationResponse represents the radar station data
type RadarStationResponse struct {
	Context    json.RawMessage        `json:"@context"`
	ID         string                 `json:"id"`
	Type       string                 `json:"type"`
	Geometry   RadarStationGeometry   `json:"geometry"`
	Properties RadarStationProperties `json:"properties"`
}

// RadarStationProperties holds detailed properties for the radar station
type RadarStationProperties struct {
	ID          string      `json:"@id"`
	Type        string      `json:"@type"`
	StationID   string      `json:"id"`
	Name        string      `json:"name"`
	StationType string      `json:"stationType"`
	Elevation   UnitValue   `json:"elevation"`
	TimeZone    string      `json:"timeZone"`
	Latency     Latency     `json:"latency"`
	RDA         RDA         `json:"rda"`
	Performance Performance `json:"performance"`
	Adaptation  Adaptation  `json:"adaptation"`
}

// Performance holds the performance-related information for the radar station
type Performance struct {
	Timestamp     string                `json:"timestamp"`
	ReportingHost string                `json:"reportingHost"`
	Properties    PerformanceProperties `json:"properties"`
}

// PerformanceProperties holds detailed performance properties
type PerformanceProperties struct {
	NtpStatus                        int       `json:"ntp_status"`
	CommandChannel                   string    `json:"commandChannel"`
	RadomeAirTemperature             UnitValue `json:"radomeAirTemperature"`
	TransitionalPowerSource          string    `json:"transitionalPowerSource"`
	HorizontalShortPulseNoise        UnitValue `json:"horizontalShortPulseNoise"`
	ElevationEncoderLight            string    `json:"elevationEncoderLight"`
	HorizontalLongPulseNoise         UnitValue `json:"horizontalLongPulseNoise"`
	AzimuthEncoderLight              string    `json:"azimuthEncoderLight"`
	HorizontalNoiseTemperature       UnitValue `json:"horizontalNoiseTemperature"`
	Linearity                        float64   `json:"linearity"`
	TransmitterPeakPower             UnitValue `json:"transmitterPeakPower"`
	HorizontalDeltadBZ0              UnitValue `json:"horizontalDeltadBZ0"`
	TransmitterRecycleCount          int       `json:"transmitterRecycleCount"`
	VerticalDeltadBZ0                UnitValue `json:"verticalDeltadBZ0"`
	ReceiverBias                     UnitValue `json:"receiverBias"`
	ShortPulseHorizontaldBZ0         UnitValue `json:"shortPulseHorizontaldBZ0"`
	TransmitterImbalance             UnitValue `json:"transmitterImbalance"`
	LongPulseHorizontaldBZ0          UnitValue `json:"longPulseHorizontaldBZ0"`
	PerformanceCheckTime             string    `json:"performanceCheckTime"`
	TransmitterLeavingAirTemperature UnitValue `json:"transmitterLeavingAirTemperature"`
	ShelterTemperature               UnitValue `json:"shelterTemperature"`
	PowerSource                      string    `json:"powerSource"`
	DynamicRange                     UnitValue `json:"dynamicRange"`
	FuelLevel                        UnitValue `json:"fuelLevel"`
}

// Adaptation holds the adaptation-related data
type Adaptation struct {
	Timestamp     string               `json:"timestamp"`
	ReportingHost string               `json:"reportingHost"`
	Properties    AdaptationProperties `json:"properties"`
}

// AdaptationProperties holds detailed adaptation properties
type AdaptationProperties struct {
	TransmitterFrequency                     UnitValue `json:"transmitterFrequency"`
	PathLossWG04Circulator                   UnitValue `json:"pathLossWG04Circulator"`
	AntennaGainIncludingRadome               UnitValue `json:"antennaGainIncludingRadome"`
	PathLossA6ArcDetector                    UnitValue `json:"pathLossA6ArcDetector"`
	CohoPowerAtA1J4                          UnitValue `json:"cohoPowerAtA1J4"`
	AmeHorizontalTestSignalPower             UnitValue `json:"ameHorzizontalTestSignalPower"`
	PathLossTransmitterCouplerCoupling       UnitValue `json:"pathLossTransmitterCouplerCoupling"`
	StaloPowerAtA1J2                         UnitValue `json:"staloPowerAtA1J2"`
	AmeNoiseSourceHorizontalExcessNoiseRatio UnitValue `json:"ameNoiseSourceHorizontalExcessNoiseRatio"`
	PathLossVerticalIFHeliaxTo4AT16          UnitValue `json:"pathLossVerticalIFHeliaxTo4AT16"`
	PathLossAT4Attenuator                    UnitValue `json:"pathLossAT4Attenuator"`
	PathLossHorzontalIFHeliaxTo4AT17         UnitValue `json:"pathLossHorzontalIFHeliaxTo4AT17"`
	PathLossIFDRIFAntiAliasFilter            UnitValue `json:"pathLossIFDRIFAntiAliasFilter"`
	PathLossIFDBurstAntiAliasFilter          UnitValue `json:"pathLossIFDBurstAntiAliasFilter"`
	PathLossWG02HarmonicFilter               UnitValue `json:"pathLossWG02HarmonicFilter"`
	TransmitterPowerDataWattsFactor          UnitValue `json:"transmitterPowerDataWattsFactor"`
	PathLossWaveguideKlystronToSwitch        UnitValue `json:"pathLossWaveguideKlystronToSwitch"`
	PulseWidthTransmitterOutputShortPulse    UnitValue `json:"pulseWidthTransmitterOutputShortPulse"`
	PulseWidthTransmitterOutputLongPulse     UnitValue `json:"pulseWidthTransmitterOutputLongPulse"`
	PathLossWG06SpectrumFilter               UnitValue `json:"pathLossWG06SpectrumFilter"`
	HorizontalReceiverNoiseShortPulse        UnitValue `json:"horizontalReceiverNoiseShortPulse"`
	HorizontalReceiverNoiseLongPulse         UnitValue `json:"horizontalReceiverNoiseLongPulse"`
	TransmitterSpectrumFilterInstalled       string    `json:"transmitterSpectrumFilterInstalled"`
}

// api.weather.gov//radar/stations/{stationId}/alarms
// RadarStationAlarmResponse represents the response containing radar station alarms
type RadarStationAlarmResponse struct {
	Context json.RawMessage     `json:"@context"`
	ID      string              `json:"@id"`
	Graph   []RadarStationAlarm `json:"@graph"`
}

// RadarStationAlarm represents a single radar station alarm
type RadarStationAlarm struct {
	Type          string `json:"@type"`
	StationID     string `json:"stationId"`
	Status        string `json:"status"`
	ActiveChannel int    `json:"activeChannel"`
	Message       string `json:"message"`
	Timestamp     string `json:"timestamp"`
}

// Common radar station response elements
// RadarStationGeometry holds the geometry data for the radar station
type RadarStationGeometry struct {
	Type        string    `json:"type"`
	Coordinates []float64 `json:"coordinates"`
}

// Elevation holds the elevation details of a radar station
type Elevation struct {
	UnitCode string  `json:"unitCode"`
	Value    float64 `json:"value"`
}

// Latency represents the latency details of a radar station
type Latency struct {
	Current                  Measurement `json:"current"`
	Average                  Measurement `json:"average"`
	Max                      Measurement `json:"max"`
	LevelTwoLastReceivedTime string      `json:"levelTwoLastReceivedTime"`
	MaxLatencyTime           string      `json:"maxLatencyTime"`
	ReportingHost            string      `json:"reportingHost"`
	Host                     string      `json:"host"`
}

// Measurement represents a measurement with a unit code and value
type Measurement struct {
	UnitCode string  `json:"unitCode"`
	Value    float64 `json:"value"`
}

// RDA represents the RDA details of a radar station
type RDA struct {
	Timestamp     string        `json:"timestamp"`
	ReportingHost string        `json:"reportingHost"`
	Properties    RDAProperties `json:"properties"`
}

// RDAProperties represents the properties of the RDA
type RDAProperties struct {
	ResolutionVersion                 *string   `json:"resolutionVersion"`
	Nl2Path                           string    `json:"nl2Path"`
	VolumeCoveragePattern             string    `json:"volumeCoveragePattern"`
	ControlStatus                     string    `json:"controlStatus"`
	BuildNumber                       float64   `json:"buildNumber"`
	AlarmSummary                      string    `json:"alarmSummary"`
	Mode                              string    `json:"mode"`
	GeneratorState                    string    `json:"generatorState"`
	SuperResolutionStatus             string    `json:"superResolutionStatus"`
	OperabilityStatus                 string    `json:"operabilityStatus"`
	Status                            string    `json:"status"`
	AverageTransmitterPower           UnitValue `json:"averageTransmitterPower"`
	ReflectivityCalibrationCorrection UnitValue `json:"reflectivityCalibrationCorrection"`
}

// UnitValue holds a value with its unit
type UnitValue struct {
	UnitCode string  `json:"unitCode"`
	Value    float64 `json:"value"`
}
