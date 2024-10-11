package nwsgo

// RadarStationResponse represents the radar station details.
type RadarStationResponse struct {
	URL         string      `json:"@id"`
	Geometry    string      `json:"geometry"`
	Name        string      `json:"name"`
	TimeZone    string      `json:"timeZone"`
	RDA         RDA         `json:"rda"`
	Performance Performance `json:"performance"`
	Adaptation  Adaptation  `json:"adaptation"`
}

type RDA struct {
	Timestamp     string        `json:"timestamp"`
	ReportingHost string        `json:"reportingHost"`
	Properties    RDAProperties `json:"properties"`
}

type RDAProperties struct {
	VolumeCoveragePattern             string    `json:"volumeCoveragePattern"`
	ResolutionVersion                 *string   `json:"resolutionVersion"` // Nullable field
	Nl2Path                           string    `json:"nl2Path"`
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

type Performance struct {
	Timestamp     string                `json:"timestamp"`
	ReportingHost string                `json:"reportingHost"`
	Properties    PerformanceProperties `json:"properties"`
}

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

type Adaptation struct {
	Timestamp     string               `json:"timestamp"`
	ReportingHost string               `json:"reportingHost"`
	Properties    AdaptationProperties `json:"properties"`
}

type AdaptationProperties struct {
	TransmitterFrequency                     UnitValue `json:"transmitterFrequency"`
	PathLossWG04Circulator                   UnitValue `json:"pathLossWG04Circulator"`
	AntennaGainIncludingRadome               UnitValue `json:"antennaGainIncludingRadome"`
	PathLossA6ArcDetector                    UnitValue `json:"pathLossA6ArcDetector"`
	CohoPowerAtA1J4                          UnitValue `json:"cohoPowerAtA1J4"`
	AmeHorzizontalTestSignalPower            UnitValue `json:"ameHorzizontalTestSignalPower"`
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

type UnitValue struct {
	UnitCode string  `json:"unitCode"`
	Value    float64 `json:"value"`
}
