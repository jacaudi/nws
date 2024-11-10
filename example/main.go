package main

import (
	"fmt"
	"log"

	"github.com/jacaudi/nwsgo"
)

var debug = false
var radarMode = ""

func main() {
	// Define the Lat & Lon
	latlon := "47.445259,-122.294533"

	pointData, err := nwsgo.GetPoints(latlon)
	if err != nil {
		log.Fatalf("Failed to get data from GPS location: %v", err)
	}

	stationID := pointData.RadarStation

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
	name := radarStation.Name
	version := radarStation.Context.Version
	Status := radarStation.RDA.Properties.Mode

	if VCP == "R35" {
		radarMode = "Clear Air Mode"
	} else if VCP == "R215" {
		radarMode = "Precipitation Mode"
	} else {
		radarMode = "Unknown Mode -- Please Update Code"
	}

	// Print the VolumeCoveragePattern
	fmt.Printf("Radar Site: %s - %s\n", stationID, name)
	fmt.Printf("Volume Coverage Pattern: %s - %s\n", VCP, radarMode)
	fmt.Printf("Status: %s\n", Status)
	if debug {
		fmt.Printf("Version: %s\n", version)
	}
}
