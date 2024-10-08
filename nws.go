// Package nws implements a basic wrapper around api.weather.gov to
// grab HTTP responses to endpoints (i.e.: weather & forecast data)
// by the National Weather Service, an agency of the United States.
package nws-go

// RadarStationList returns the list of radar stations.
func RadarStationList() (*RadarStationListResponse, error) {
    endpoint := config.endpointRadarStations()

    radarStations := &RadarStationListResponse{}
    if err := config.GetAndDecode(endpoint, radarStations); err != nil {
        return nil, err
    }

    return radarStations, nil
}

// RadarStation returns details for a single radar station by stationId.
func RadarStation(stationID string) (*RadarStationResponse, error) {
    endpoint := config.endpointRadarStation(stationID)

    radarStation := &RadarStationResponse{}
    if err := config.GetAndDecode(endpoint, radarStation); err != nil {
        return nil, err
    }

    return radarStation, nil
}

// RadarStationAlarms returns alarm details for a radar station by stationId.
func RadarStationAlarms(stationID string) (*RadarStationAlarmResponse, error) {
    endpoint := config.endpointRadarStationAlarms(stationID)

    radarAlarms := &RadarStationAlarmResponse{}
    if err := config.GetAndDecode(endpoint, radarAlarms); err != nil {
        return nil, err
    }

    return radarAlarms, nil
}
