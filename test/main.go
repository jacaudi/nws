package main

import (
	"fmt"
	"log"

	"github.com/jacaudi/nwsgo"
)

func main() {
	// Step 1: Call nwsgo.RadarStationList() to get the RadarStationListResponse
	radarStationListResponse, err := nwsgo.RadarStationList()
	if err != nil {
		log.Fatalf("Error fetching radar station list: %v", err)
	}

	fmt.Printf("RadarStationListResponse: %+v\n", radarStationListResponse)

	// Step 2: Iterate over each RadarStationListFeature in the response
	for _, feature := range radarStationListResponse.Features {
		stationID := feature.Properties.StationID
		volumeCoveragePattern := feature.Properties.RDA.Properties.VolumeCoveragePattern

		// Step 3: Store or process the VolumeCoveragePattern as needed
		fmt.Printf("StationID: %s, VolumeCoveragePattern: %+v\n", stationID, volumeCoveragePattern)
	}
}
