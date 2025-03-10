package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/jacaudi/nws/cmd/nws"
)

var (
	debug, _  = strconv.ParseBool(os.Getenv("DEBUG"))
	radarMode = ""
)

func main() {
	// Define the Lat & Lon
	latlon := "47.445259,-122.294533"

	pointData, err := nws.GetPoints(latlon)
	if err != nil {
		log.Fatalf("Failed to get data from GPS location: %v", err)
	}

	if debug {
		log.Printf("Points Endpoint Response: %v\n\n", pointData)
	}

	stationID := pointData.RadarStation

	// Get the radar station details for KATX
	radarStation, err := nws.RadarStation(stationID)
	if err != nil {
		log.Fatalf("Failed to get radar station details: %v", err)
	}

	// Print the entire radarStation object for debugging
	if debug {
		fmt.Printf("RadarStation details: %+v\n\n", radarStation)
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
