package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/jacaudi/nws/cmd/nws"
)

var debug, _ = strconv.ParseBool(os.Getenv("DEBUG"))

func main() {
	client, err := nws.NewClient(
		nws.WithUserAgent("nws-example-radar/1.0 (+https://github.com/jacaudi/nws)"),
	)
	if err != nil {
		log.Fatalf("NewClient: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	latlon := "47.445259,-122.294533"
	pointData, err := client.GetPoints(ctx, latlon)
	if err != nil {
		log.Fatalf("GetPoints: %v", err)
	}
	if debug {
		log.Printf("Points: %+v", pointData)
	}

	stationID := pointData.RadarStation
	station, err := client.RadarStation(ctx, stationID)
	if err != nil {
		log.Fatalf("RadarStation: %v", err)
	}
	if debug {
		fmt.Printf("RadarStation: %+v\n\n", station)
	}

	vcp := station.RDA.Properties.VolumeCoveragePattern
	mode := "Unknown Mode -- Please Update Code"
	switch vcp {
	case "R35":
		mode = "Clear Air Mode"
	case "R215":
		mode = "Precipitation Mode"
	}

	fmt.Printf("Radar Site: %s - %s\n", stationID, station.Name)
	fmt.Printf("Volume Coverage Pattern: %s - %s\n", vcp, mode)
	fmt.Printf("Status: %s\n", station.RDA.Properties.Mode)
	if debug {
		fmt.Printf("Version: %s\n", station.Context.Version)
	}
}
