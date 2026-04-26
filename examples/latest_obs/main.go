package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/jacaudi/nws/cmd/nws"
)

var debug, _ = strconv.ParseBool(os.Getenv("DEBUG"))

func celsiusToFahrenheit(c float64) float64 { return (c * 9 / 5) + 32 }

func convertUnit(unitCode string) string {
	switch unitCode {
	case "wmoUnit:degC":
		return "°C"
	case "wmoUnit:km_h-1":
		return "km/h"
	case "wmoUnit:Pa":
		return "Pa"
	case "wmoUnit:m":
		return "m"
	case "wmoUnit:percent":
		return "%"
	default:
		return unitCode
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the Station ID: ")
	stationID, _ := reader.ReadString('\n')
	stationID = strings.TrimSpace(stationID)
	if stationID == "" {
		log.Fatal("Station ID cannot be empty")
	}

	client, err := nws.NewClient(
		nws.WithUserAgent("nws-example-obs/1.0 (+https://github.com/jacaudi/nws)"),
	)
	if err != nil {
		log.Fatalf("NewClient: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	obs, err := client.GetLatestObservations(ctx, stationID)
	if err != nil {
		log.Fatalf("GetLatestObservations: %v", err)
	}
	if debug {
		fmt.Printf("Observation Details: %+v\n\n", obs)
	}

	temp := obs.Temperature.Value
	tempUnit := convertUnit(obs.Temperature.UnitCode)
	tempF := celsiusToFahrenheit(temp)
	dew := obs.Dewpoint.Value
	dewUnit := convertUnit(obs.Dewpoint.UnitCode)
	dewF := celsiusToFahrenheit(dew)
	pressure := obs.BarometricPressure.Value / 100

	fmt.Printf("Station: %s\n", stationID)
	fmt.Printf("Temperature: %.1f%s - %.1f°F\n", temp, tempUnit, tempF)
	fmt.Printf("Dew Point: %.1f%s - %.1f°F\n", dew, dewUnit, dewF)
	fmt.Printf("Pressure: %.0f hPa\n", pressure)
}
