package nwsgo

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/jacaudi/nwsgo/internal/endpoints/radar"
)

func TestRadarStation(t *testing.T) {
	// Test case 1: Successful response
	stationID := "KATX"
	expectedResponse := &radar.RadarStationResponse{
		// Define the expected response here
	}
	response, err := RadarStation(stationID)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if !reflect.DeepEqual(response, expectedResponse) {
		t.Errorf("Unexpected response. Expected: %+v, Got: %+v", expectedResponse, response)
	}

	// Test case 2: Error response
	stationID = "XYZ789"
	expectedError := fmt.Errorf("failed to unmarshal response: some error")
	response, err = RadarStation(stationID)
	if err == nil {
		t.Errorf("Expected error, but got nil")
	}
	if err.Error() != expectedError.Error() {
		t.Errorf("Unexpected error. Expected: %v, Got: %v", expectedError, err)
	}
}
