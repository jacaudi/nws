package main

import (
	"fmt"
	"log"

	"github.com/jacaudi/nwsgo"
)

var debug = false

func main() {
	// Define the station ID for KATX
	stationID := "KATX"

	// Get the radar station details for KATX
	radarStation, err := nwsgo.RadarStation(stationID)
	if err != nil {
		log.Fatalf("Failed to get radar station details: %v", err)
	}

	// Print the entire radarStation object for debugging
	if debug {
		fmt.Printf("RadarStation details: %+v\n", radarStation)
	}

	// Extract the VolumeCoveragePattern from the radar station details
	volumeCoveragePattern := radarStation.RDA.Properties.VolumeCoveragePattern

	// Print the VolumeCoveragePattern
	fmt.Printf("VCP for station %s: %v\n", stationID, volumeCoveragePattern)
}
