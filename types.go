package nws

// QuantitativeValue is available for various statistics and can be
// enabled with an optional request header to the nws API. In the
// future it is expected at that QV will replace single values such
// as measurements for Temperature
type QuantitativeValue struct {
	Value          float64 `json:"value"`
	MaxValue       float64 `json:"maxValue"`
	MinValue       float64 `json:"minValue"`
	UnitCode       string  `json:"unitCode"`
	QualityControl string  `json:"qualityControl"`
}

// PointsResponse holds the JSON values from /points/<lat,lon>
type PointsResponse struct {
	ID                          string `json:"@id"`
	CWA                         string `json:"cwa"`
	Office                      string `json:"forecastOffice"`
	GridX                       int64  `json:"gridX"`
	GridY                       int64  `json:"gridY"`
	EndpointForecast            string `json:"forecast"`
	EndpointForecastHourly      string `json:"forecastHourly"`
	EndpointObservationStations string `json:"observationStations"`
	EndpointForecastGridData    string `json:"forecastGridData"`
	Timezone                    string `json:"timeZone"`
	RadarStation                string `json:"radarStation"`
}

// OfficeAddress holds the JSON values for the address of an OfficeResponse
type OfficeAddress struct {
	Type          string `json:"@type"`
	StreetAddress string `json:"streetAddress"`
	Locality      string `json:"addressLocality"`
	Region        string `json:"addressRegion"`
	PostalCode    string `json:"postalCode"`
}

// OfficeResponse holds the JSON values from /offices/<id>
type OfficeResponse struct {
	Type                        string        `json:"@type"`
	URI                         string        `json:"@id"`
	ID                          string        `json:"id"`
	Name                        string        `json:"name"`
	Address                     OfficeAddress `json:"address"`
	Telephone                   string        `json:"telephone"`
	FaxNumber                   string        `json:"faxNumber"`
	Email                       string        `json:"email"`
	SameAs                      string        `json:"sameAs"`
	NWSRegion                   string        `json:"nwsRegion"`
	ParentOrganization          string        `json:"parentOrganization"`
	ResponsibleCounties         []string      `json:"responsibleCounties"`
	ResponsibleForecastZones    []string      `json:"responsibleForecastZones"`
	ResponsibleFireZones        []string      `json:"responsibleFireZones"`
	ApprovedObservationStations []string      `json:"approvedObservationStations"`
}

// StationsResponse holds the JSON values from /points/<lat,lon>/stations
type StationsResponse struct {
	Stations []string `json:"observationStations"`
}

// ForecastElevation holds the JSON values for a forecast response's elevation.
type ForecastElevation struct {
	Value float64 `json:"value"`
	Units string  `json:"unitCode"`
}

// ForecastResponsePeriod holds the JSON values for a period within a forecast response.
type ForecastResponsePeriod struct {
	ID               int32   `json:"number"`
	Name             string  `json:"name"`
	StartTime        string  `json:"startTime"`
	EndTime          string  `json:"endTime"`
	IsDaytime        bool    `json:"isDaytime"`
	Temperature      float64 // preserved for legacy compatibility, may be deprecated in the future
	TemperatureUnit  string  // preserved for legacy compatibility, may be deprecated in the future
	TemperatureTrend string  `json:"temperatureTrend"`
	WindSpeed        string  // preserved for legacy compatibility, may be deprecated in the future
	WindDirection    string  `json:"windDirection"`
	Icon             string  `json:"icon"`
	Summary          string  `json:"shortForecast"`
	Details          string  `json:"detailedForecast"`

	QuantitativeProbability      QuantitativeValue `json:"probabilityOfPrecipitation"`
	QuantitativeDewpoint         QuantitativeValue `json:"dewpoint"`
	QuantitativeRelativeHumidity QuantitativeValue `json:"relativeHumidity"`
	QuantitativeTemperature      QuantitativeValue `json:"temperature"`
	QuantitativeWindSpeed        QuantitativeValue `json:"windSpeed"`
	QuantitativeWindGust         QuantitativeValue `json:"windGust"`
}

// ForecastResponsePeriodHourly provides the JSON value for a period within an hourly forecast.
type ForecastResponsePeriodHourly = ForecastResponsePeriod

// ForecastResponse holds the JSON values from /gridpoints/<cwa>/<x,y>/forecast"
type ForecastResponse struct {
	Updated   string                   `json:"updated"`
	Units     string                   `json:"units"`
	Elevation ForecastElevation        `json:"elevation"`
	Periods   []ForecastResponsePeriod `json:"periods"`
	Point     *PointsResponse
}

// WeatherValueItem holds the JSON values for a weather.values[x].value.
type WeatherValueItem struct {
	Coverage  string `json:"coverage"`
	Weather   string `json:"weather"`
	Intensity string `json:"intensity"`
}

// WeatherValue holds the JSON value for a weather.values[x] value.
type WeatherValue struct {
	ValidTime string             `json:"validTime"` // ISO 8601 time interval, e.g. 2019-07-04T18:00:00+00:00/PT3H
	Value     []WeatherValueItem `json:"value"`
}

// Weather holds the JSON value for the weather object.
type Weather struct {
	Values []WeatherValue `json:"values"`
}

// HazardValueItem holds a value item from a GridpointForecastResponse's
// hazard.values[x].value[x].
type HazardValueItem struct {
	Phenomenon   string `json:"phenomenon"`
	Significance string `json:"significance"`
	EventNumber  int32  `json:"event_number"`
}

// HazardValue holds a hazard value from a GridpointForecastResponse's
// hazard.values[x].
type HazardValue struct {
	ValidTime string            `json:"validTime"` // ISO 8601 time interval, e.g. 2019-07-04T18:00:00+00:00/PT3H
	Value     []HazardValueItem `json:"value"`
}

// Hazard holds a slice of HazardValue items from a GridpointForecastResponse hazards
type Hazard struct {
	Values []HazardValue `json:"values"`
}

// HourlyForecastResponse holds the JSON values for the hourly forecast.
type HourlyForecastResponse struct {
	Updated           string                         `json:"updated"`
	Units             string                         `json:"units"`
	ForecastGenerator string                         `json:"forecastGenerator"`
	GeneratedAt       string                         `json:"generatedAt"`
	UpdateTime        string                         `json:"updateTime"`
	ValidTimes        string                         `json:"validTimes"`
	Periods           []ForecastResponsePeriodHourly `json:"periods"`
	Point             *PointsResponse
}

// GridpointForecastResponse holds the JSON values from /gridpoints/<cwa>/<x,y>"
// See https://weather-gov.github.io/api/gridpoints for information.
type GridpointForecastResponse struct {
	Updated                          string                      `json:"updateTime"`
	Elevation                        ForecastElevation           `json:"elevation"`
	Weather                          Weather                     `json:"weather"`
	Hazards                          Hazard                      `json:"hazards"`
	Temperature                      GridpointForecastTimeSeries `json:"temperature"`
	Dewpoint                         GridpointForecastTimeSeries `json:"dewpoint"`
	MaxTemperature                   GridpointForecastTimeSeries `json:"maxTemperature"`
	MinTemperature                   GridpointForecastTimeSeries `json:"minTemperature"`
	RelativeHumidity                 GridpointForecastTimeSeries `json:"relativeHumidity"`
	ApparentTemperature              GridpointForecastTimeSeries `json:"apparentTemperature"`
	HeatIndex                        GridpointForecastTimeSeries `json:"heatIndex"`
	WindChill                        GridpointForecastTimeSeries `json:"windChill"`
	SkyCover                         GridpointForecastTimeSeries `json:"skyCover"`
	WindDirection                    GridpointForecastTimeSeries `json:"windDirection"`
	WindSpeed                        GridpointForecastTimeSeries `json:"windSpeed"`
	WindGust                         GridpointForecastTimeSeries `json:"windGust"`
	ProbabilityOfPrecipitation       GridpointForecastTimeSeries `json:"probabilityOfPrecipitation"`
	QuantitativePrecipitation        GridpointForecastTimeSeries `json:"quantitativePrecipitation"`
	IceAccumulation                  GridpointForecastTimeSeries `json:"iceAccumulation"`
	SnowfallAmount                   GridpointForecastTimeSeries `json:"snowfallAmount"`
	SnowLevel                        GridpointForecastTimeSeries `json:"snowLevel"`
	CeilingHeight                    GridpointForecastTimeSeries `json:"ceilingHeight"`
	Visibility                       GridpointForecastTimeSeries `json:"visibility"`
	TransportWindSpeed               GridpointForecastTimeSeries `json:"transportWindSpeed"`
	TransportWindDirection           GridpointForecastTimeSeries `json:"transportWindDirection"`
	MixingHeight                     GridpointForecastTimeSeries `json:"mixingHeight"`
	HainesIndex                      GridpointForecastTimeSeries `json:"hainesIndex"`
	LightningActivityLevel           GridpointForecastTimeSeries `json:"lightningActivityLevel"`
	TwentyFootWindSpeed              GridpointForecastTimeSeries `json:"twentyFootWindSpeed"`
	TwentyFootWindDirection          GridpointForecastTimeSeries `json:"twentyFootWindDirection"`
	WaveHeight                       GridpointForecastTimeSeries `json:"waveHeight"`
	WavePeriod                       GridpointForecastTimeSeries `json:"wavePeriod"`
	WaveDirection                    GridpointForecastTimeSeries `json:"waveDirection"`
	PrimarySwellHeight               GridpointForecastTimeSeries `json:"primarySwellHeight"`
	PrimarySwellDirection            GridpointForecastTimeSeries `json:"primarySwellDirection"`
	SecondarySwellHeight             GridpointForecastTimeSeries `json:"secondarySwellHeight"`
	SecondarySwellDirection          GridpointForecastTimeSeries `json:"secondarySwellDirection"`
	WavePeriod2                      GridpointForecastTimeSeries `json:"wavePeriod2"`
	WindWaveHeight                   GridpointForecastTimeSeries `json:"windWaveHeight"`
	DispersionIndex                  GridpointForecastTimeSeries `json:"dispersionIndex"`
	Pressure                         GridpointForecastTimeSeries `json:"pressure"`
	ProbabilityOfTropicalStormWinds  GridpointForecastTimeSeries `json:"probabilityOfTropicalStormWinds"`
	ProbabilityOfHurricaneWinds      GridpointForecastTimeSeries `json:"probabilityOfHurricaneWinds"`
	PotentialOf15mphWinds            GridpointForecastTimeSeries `json:"potentialOf15mphWinds"`
	PotentialOf25mphWinds            GridpointForecastTimeSeries `json:"potentialOf25mphWinds"`
	PotentialOf35mphWinds            GridpointForecastTimeSeries `json:"potentialOf35mphWinds"`
	PotentialOf45mphWinds            GridpointForecastTimeSeries `json:"potentialOf45mphWinds"`
	PotentialOf20mphWindGusts        GridpointForecastTimeSeries `json:"potentialOf20mphWindGusts"`
	PotentialOf30mphWindGusts        GridpointForecastTimeSeries `json:"potentialOf30mphWindGusts"`
	PotentialOf40mphWindGusts        GridpointForecastTimeSeries `json:"potentialOf40mphWindGusts"`
	PotentialOf50mphWindGusts        GridpointForecastTimeSeries `json:"potentialOf50mphWindGusts"`
	PotentialOf60mphWindGusts        GridpointForecastTimeSeries `json:"potentialOf60mphWindGusts"`
	GrasslandFireDangerIndex         GridpointForecastTimeSeries `json:"grasslandFireDangerIndex"`
	ProbabilityOfThunder             GridpointForecastTimeSeries `json:"probabilityOfThunder"`
	DavisStabilityIndex              GridpointForecastTimeSeries `json:"davisStabilityIndex"`
	AtmosphericDispersionIndex       GridpointForecastTimeSeries `json:"atmosphericDispersionIndex"`
	LowVisibilityOccurrenceRiskIndex GridpointForecastTimeSeries `json:"lowVisibilityOccurrenceRiskIndex"`
	Stability                        GridpointForecastTimeSeries `json:"stability"`
	RedFlagThreatIndex               GridpointForecastTimeSeries `json:"redFlagThreatIndex"`
	Point                            *PointsResponse
}

// api.weather.gov/radar/stations
// RadarStationListResponse holds the JSON values from /radar/stations
type RadarStationListResponse struct {
	Context  []interface{} `json:"@context"`
	Type     string        `json:"type"`
	Features []RadarStationListFeature `json:"features"`
}

// RadarStationListFeature represents an individual radar station feature
type RadarStationListFeature struct {
	ID         string              `json:"id"`
	Type       string              `json:"type"`
	Geometry   RadarStationGeometry `json:"geometry"`
	Properties RadarStationListProperties `json:"properties"`
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

// api.weather.gov/radar/stations/{stationId}
// RadarStationResponse represents the radar station data
type RadarStationResponse struct {
	Context    []interface{}          `json:"@context"`
	ID         string                 `json:"id"`
	Type       string                 `json:"type"`
	Geometry   RadarStationGeometry   `json:"geometry"`
	Properties RadarStationProperties `json:"properties"`
}

// RadarStationProperties holds detailed properties for the radar station
type RadarStationProperties struct {
	ID                           string                 `json:"@id"`
	Type                         string                 `json:"@type"`
	StationID                    string                 `json:"id"`
	Name                         string                 `json:"name"`
	StationType                  string                 `json:"stationType"`
	Elevation                    UnitValue              `json:"elevation"`
	TimeZone                     string                 `json:"timeZone"`
	Latency                      Latency                `json:"latency"`
	RDA                          RDA                    `json:"rda"`
	Performance                  Performance            `json:"performance"`
	Adaptation                   Adaptation             `json:"adaptation"`
}

// Performance holds the performance-related information for the radar station
type Performance struct {
	Timestamp     string                `json:"timestamp"`
	ReportingHost string                `json:"reportingHost"`
	Properties    PerformanceProperties `json:"properties"`
}

// PerformanceProperties holds detailed performance properties
type PerformanceProperties struct {
	NtpStatus                   int       `json:"ntp_status"`
	CommandChannel              string    `json:"commandChannel"`
	RadomeAirTemperature        UnitValue `json:"radomeAirTemperature"`
	TransitionalPowerSource      string    `json:"transitionalPowerSource"`
	HorizontalShortPulseNoise   UnitValue `json:"horizontalShortPulseNoise"`
	ElevationEncoderLight       string    `json:"elevationEncoderLight"`
	HorizontalLongPulseNoise    UnitValue `json:"horizontalLongPulseNoise"`
	AzimuthEncoderLight         string    `json:"azimuthEncoderLight"`
	HorizontalNoiseTemperature  UnitValue `json:"horizontalNoiseTemperature"`
	Linearity                   float64   `json:"linearity"`
	TransmitterPeakPower        UnitValue `json:"transmitterPeakPower"`
	HorizontalDeltadBZ0         UnitValue `json:"horizontalDeltadBZ0"`
	TransmitterRecycleCount     int       `json:"transmitterRecycleCount"`
	VerticalDeltadBZ0           UnitValue `json:"verticalDeltadBZ0"`
	ReceiverBias                UnitValue `json:"receiverBias"`
	ShortPulseHorizontaldBZ0    UnitValue `json:"shortPulseHorizontaldBZ0"`
	TransmitterImbalance        UnitValue `json:"transmitterImbalance"`
	LongPulseHorizontaldBZ0     UnitValue `json:"longPulseHorizontaldBZ0"`
	PerformanceCheckTime        string    `json:"performanceCheckTime"`
	TransmitterLeavingAirTemperature UnitValue `json:"transmitterLeavingAirTemperature"`
	ShelterTemperature          UnitValue `json:"shelterTemperature"`
	PowerSource                 string    `json:"powerSource"`
	DynamicRange                UnitValue `json:"dynamicRange"`
	FuelLevel                   UnitValue `json:"fuelLevel"`
}

// Adaptation holds the adaptation-related data
type Adaptation struct {
	Timestamp     string               `json:"timestamp"`
	ReportingHost string               `json:"reportingHost"`
	Properties    AdaptationProperties `json:"properties"`
}

// AdaptationProperties holds detailed adaptation properties
type AdaptationProperties struct {
	TransmitterFrequency                      UnitValue `json:"transmitterFrequency"`
	PathLossWG04Circulator                    UnitValue `json:"pathLossWG04Circulator"`
	AntennaGainIncludingRadome                UnitValue `json:"antennaGainIncludingRadome"`
	PathLossA6ArcDetector                     UnitValue `json:"pathLossA6ArcDetector"`
	CohoPowerAtA1J4                           UnitValue `json:"cohoPowerAtA1J4"`
	AmeHorizontalTestSignalPower              UnitValue `json:"ameHorzizontalTestSignalPower"`
	PathLossTransmitterCouplerCoupling        UnitValue `json:"pathLossTransmitterCouplerCoupling"`
	StaloPowerAtA1J2                          UnitValue `json:"staloPowerAtA1J2"`
	AmeNoiseSourceHorizontalExcessNoiseRatio  UnitValue `json:"ameNoiseSourceHorizontalExcessNoiseRatio"`
	PathLossVerticalIFHeliaxTo4AT16           UnitValue `json:"pathLossVerticalIFHeliaxTo4AT16"`
	PathLossAT4Attenuator                     UnitValue `json:"pathLossAT4Attenuator"`
	PathLossHorzontalIFHeliaxTo4AT17          UnitValue `json:"pathLossHorzontalIFHeliaxTo4AT17"`
	PathLossIFDRIFAntiAliasFilter             UnitValue `json:"pathLossIFDRIFAntiAliasFilter"`
	PathLossIFDBurstAntiAliasFilter           UnitValue `json:"pathLossIFDBurstAntiAliasFilter"`
	PathLossWG02HarmonicFilter                UnitValue `json:"pathLossWG02HarmonicFilter"`
	TransmitterPowerDataWattsFactor           UnitValue `json:"transmitterPowerDataWattsFactor"`
	PathLossWaveguideKlystronToSwitch         UnitValue `json:"pathLossWaveguideKlystronToSwitch"`
	PulseWidthTransmitterOutputShortPulse     UnitValue `json:"pulseWidthTransmitterOutputShortPulse"`
	PulseWidthTransmitterOutputLongPulse      UnitValue `json:"pulseWidthTransmitterOutputLongPulse"`
	PathLossWG06SpectrumFilter                UnitValue `json:"pathLossWG06SpectrumFilter"`
	HorizontalReceiverNoiseShortPulse         UnitValue `json:"horizontalReceiverNoiseShortPulse"`
	HorizontalReceiverNoiseLongPulse          UnitValue `json:"horizontalReceiverNoiseLongPulse"`
	TransmitterSpectrumFilterInstalled        string    `json:"transmitterSpectrumFilterInstalled"`
}

// api.weather.gov//radar/stations/{stationId}/alarms
// RadarStationAlarmResponse represents the response containing radar station alarms
type RadarStationAlarmResponse struct {
	Context  []interface{}       `json:"@context"`
	ID       string              `json:"@id"`
	Graph    []RadarStationAlarm `json:"@graph"`
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

// Latency represents the latency-related information of the radar station
type Latency struct {
	Current                  UnitValue `json:"current"`
	Average                  UnitValue `json:"average"`
	Max                      UnitValue `json:"max"`
	LevelTwoLastReceivedTime string    `json:"levelTwoLastReceivedTime"`
	MaxLatencyTime           string    `json:"maxLatencyTime"`
	ReportingHost            string    `json:"reportingHost"`
	Host                     string    `json:"host"`
}

// RDA holds the Radar Data Acquisition (RDA) information of the radar station
type RDA struct {
	Timestamp     string       `json:"timestamp"`
	ReportingHost string       `json:"reportingHost"`
	Properties    RDAProperties `json:"properties"`
}

// RDAProperties holds various operational properties of the RDA
type RDAProperties struct {
	ResolutionVersion                 *string    `json:"resolutionVersion"`
	Nl2Path                           string     `json:"nl2Path"`
	VolumeCoveragePattern             string     `json:"volumeCoveragePattern"`
	ControlStatus                     string     `json:"controlStatus"`
	BuildNumber                       float64    `json:"buildNumber"`
	AlarmSummary                      string     `json:"alarmSummary"`
	Mode                              string     `json:"mode"`
	GeneratorState                    string     `json:"generatorState"`
	SuperResolutionStatus             string     `json:"superResolutionStatus"`
	OperabilityStatus                 string     `json:"operabilityStatus"`
	Status                            string     `json:"status"`
	AverageTransmitterPower           UnitValue  `json:"averageTransmitterPower"`
	ReflectivityCalibrationCorrection UnitValue `json:"reflectivityCalibrationCorrection"`
}

// UnitValue holds a value with its unit
type UnitValue struct {
	UnitCode string  `json:"unitCode"`
	Value    float64 `json:"value"`
}

// GridpointForecastTimeSeriesValue holds the JSON value for a
// GridpointForecastTimeSeries' values[x] item.
type GridpointForecastTimeSeriesValue struct {
	ValidTime string  `json:"validTime"` // ISO 8601 time interval, e.g. 2019-07-04T18:00:00+00:00/PT3H
	Value     float64 `json:"value"`
}

// GridpointForecastTimeSeries holds a series of data from a gridpoint forecast
type GridpointForecastTimeSeries struct {
	Uom    string                             `json:"uom"` // Unit of Measure
	Values []GridpointForecastTimeSeriesValue `json:"values"`
}
