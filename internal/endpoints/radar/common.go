package radar

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
)

// StringOrNumber accepts a JSON value that may be encoded as either a string
// or a number. The NWS API is inconsistent across stations for some fields
// (e.g. resolutionVersion), so values are normalized to their string form.
type StringOrNumber string

func (s *StringOrNumber) UnmarshalJSON(data []byte) error {
	if len(data) == 0 || bytes.Equal(data, []byte("null")) {
		*s = ""
		return nil
	}
	if data[0] == '"' {
		var str string
		if err := json.Unmarshal(data, &str); err != nil {
			return err
		}
		*s = StringOrNumber(str)
		return nil
	}
	var n json.Number
	dec := json.NewDecoder(bytes.NewReader(data))
	dec.UseNumber()
	if err := dec.Decode(&n); err != nil {
		return fmt.Errorf("StringOrNumber: %w", err)
	}
	*s = StringOrNumber(n.String())
	return nil
}

func (s StringOrNumber) MarshalJSON() ([]byte, error) {
	if s == "" {
		return []byte("null"), nil
	}
	if _, err := strconv.ParseFloat(string(s), 64); err == nil {
		return []byte(s), nil
	}
	return json.Marshal(string(s))
}

type Adaptation struct {
	Properties    AdaptationProperties `json:"properties"`
	ReportingHost string               `json:"reportingHost"`
	Timestamp     string               `json:"timestamp"`
}

type AdaptationProperties struct {
	AmeHorzizontalTestSignalPower            UnitValue `json:"ameHorzizontalTestSignalPower"`
	AmeNoiseSourceHorizontalExcessNoiseRatio UnitValue `json:"ameNoiseSourceHorizontalExcessNoiseRatio"`
	AntennaGainIncludingRadome               UnitValue `json:"antennaGainIncludingRadome"`
	CohoPowerAtA1J4                          UnitValue `json:"cohoPowerAtA1J4"`
	HorizontalReceiverNoiseLongPulse         UnitValue `json:"horizontalReceiverNoiseLongPulse"`
	HorizontalReceiverNoiseShortPulse        UnitValue `json:"horizontalReceiverNoiseShortPulse"`
	PathLossA6ArcDetector                    UnitValue `json:"pathLossA6ArcDetector"`
	PathLossAT4Attenuator                    UnitValue `json:"pathLossAT4Attenuator"`
	PathLossHorzontalIFHeliaxTo4AT17         UnitValue `json:"pathLossHorzontalIFHeliaxTo4AT17"`
	PathLossIFDBurstAntiAliasFilter          UnitValue `json:"pathLossIFDBurstAntiAliasFilter"`
	PathLossIFDRIFAntiAliasFilter            UnitValue `json:"pathLossIFDRIFAntiAliasFilter"`
	PathLossTransmitterCouplerCoupling       UnitValue `json:"pathLossTransmitterCouplerCoupling"`
	PathLossVerticalIFHeliaxTo4AT16          UnitValue `json:"pathLossVerticalIFHeliaxTo4AT16"`
	PathLossWaveguideKlystronToSwitch        UnitValue `json:"pathLossWaveguideKlystronToSwitch"`
	PathLossWG02HarmonicFilter               UnitValue `json:"pathLossWG02HarmonicFilter"`
	PathLossWG04Circulator                   UnitValue `json:"pathLossWG04Circulator"`
	PathLossWG06SpectrumFilter               UnitValue `json:"pathLossWG06SpectrumFilter"`
	PulseWidthTransmitterOutputLongPulse     UnitValue `json:"pulseWidthTransmitterOutputLongPulse"`
	PulseWidthTransmitterOutputShortPulse    UnitValue `json:"pulseWidthTransmitterOutputShortPulse"`
	StaloPowerAtA1J2                         UnitValue `json:"staloPowerAtA1J2"`
	TransmitterFrequency                     UnitValue `json:"transmitterFrequency"`
	TransmitterPowerDataWattsFactor          UnitValue `json:"transmitterPowerDataWattsFactor"`
	TransmitterSpectrumFilterInstalled       string    `json:"transmitterSpectrumFilterInstalled"`
}

type Context struct {
	NWSUnit string `json:"nwsUnit"`
	Version string `json:"@version"`
	WMOUnit string `json:"wmoUnit"`
}

type Latency struct {
	Average                  UnitValue `json:"average"`
	Current                  UnitValue `json:"current"`
	Host                     string    `json:"host"`
	LevelTwoLastReceivedTime string    `json:"levelTwoLastReceivedTime"`
	Max                      UnitValue `json:"max"`
	MaxLatencyTime           string    `json:"maxLatencyTime"`
	ReportingHost            string    `json:"reportingHost"`
}

type Performance struct {
	Properties    PerformanceProperties `json:"properties"`
	ReportingHost string                `json:"reportingHost"`
	Timestamp     string                `json:"timestamp"`
}

type PerformanceProperties struct {
	AzimuthEncoderLight              string    `json:"azimuthEncoderLight"`
	CommandChannel                   string    `json:"commandChannel"`
	DynamicRange                     UnitValue `json:"dynamicRange"`
	ElevationEncoderLight            string    `json:"elevationEncoderLight"`
	FuelLevel                        UnitValue `json:"fuelLevel"`
	HorizontalDeltadBZ0              UnitValue `json:"horizontalDeltadBZ0"`
	HorizontalLongPulseNoise         UnitValue `json:"horizontalLongPulseNoise"`
	HorizontalNoiseTemperature       UnitValue `json:"horizontalNoiseTemperature"`
	HorizontalShortPulseNoise        UnitValue `json:"horizontalShortPulseNoise"`
	Linearity                        float64   `json:"linearity"`
	LongPulseHorizontaldBZ0          UnitValue `json:"longPulseHorizontaldBZ0"`
	NtpStatus                        int       `json:"ntp_status"`
	PerformanceCheckTime             string    `json:"performanceCheckTime"`
	PowerSource                      string    `json:"powerSource"`
	RadomeAirTemperature             UnitValue `json:"radomeAirTemperature"`
	ReceiverBias                     UnitValue `json:"receiverBias"`
	ShelterTemperature               UnitValue `json:"shelterTemperature"`
	ShortPulseHorizontaldBZ0         UnitValue `json:"shortPulseHorizontaldBZ0"`
	TransitionalPowerSource          string    `json:"transitionalPowerSource"`
	TransmitterImbalance             UnitValue `json:"transmitterImbalance"`
	TransmitterLeavingAirTemperature UnitValue `json:"transmitterLeavingAirTemperature"`
	TransmitterPeakPower             UnitValue `json:"transmitterPeakPower"`
	TransmitterRecycleCount          int       `json:"transmitterRecycleCount"`
	VerticalDeltadBZ0                UnitValue `json:"verticalDeltadBZ0"`
}

type RDA struct {
	Properties    RDAProperties `json:"properties"`
	ReportingHost string        `json:"reportingHost"`
	Timestamp     string        `json:"timestamp"`
}

type RDAProperties struct {
	AlarmSummary                      string    `json:"alarmSummary"`
	AverageTransmitterPower           UnitValue `json:"averageTransmitterPower"`
	BuildNumber                       float64   `json:"buildNumber"`
	ControlStatus                     string    `json:"controlStatus"`
	GeneratorState                    string    `json:"generatorState"`
	Mode                              string    `json:"mode"`
	Nl2Path                           string    `json:"nl2Path"`
	OperabilityStatus                 string    `json:"operabilityStatus"`
	ReflectivityCalibrationCorrection UnitValue `json:"reflectivityCalibrationCorrection"`
	ResolutionVersion                 *StringOrNumber `json:"resolutionVersion"` // Nullable; API may return string or number
	Status                            string    `json:"status"`
	SuperResolutionStatus             string    `json:"superResolutionStatus"`
	VolumeCoveragePattern             string    `json:"volumeCoveragePattern"`
}

type UnitValue struct {
	UnitCode string  `json:"unitCode"`
	Value    float64 `json:"value"`
}
