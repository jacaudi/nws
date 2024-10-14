package main

import (
	"fmt"
	"log"

	"github.com/jacaudi/nwsgo"
)

var debug = false
var radarMode = ""

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

	// Extract the Values from the radar station details
	VCP := radarStation.RDA.Properties.VolumeCoveragePattern
	version := radarStation.Context.Version
	Status := radarStation.RDA.Properties.Mode

	if VCP == "R35" {
		radarMode = "Clear Air Mode"
	} else if VCP == "R215" {
		radarMode = "Precipatation Mode"
	} else {
		radarMode = "Unknown Mode -- Please Update Code"
	}

	// Print the VolumeCoveragePattern
	fmt.Printf("Radar Site: %s\n", stationID)
	fmt.Printf("Radar Mode: %s\n", radarMode)
	fmt.Printf("Radar Status: %s\n", Status)
	if debug {
		fmt.Printf("Version: %s\n", version)
	}
}
