package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/jacaudi/nws/cmd/nws"
)

var (
	debug, _ = strconv.ParseBool(os.Getenv("DEBUG"))
)

func celsiusToFahrenheit(celsius float64) float64 {
	return (celsius * 9 / 5) + 32
}

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

// It retrieves the latest observations for a given station ID,
// converts the temperature and dew point values to Fahrenheit,
// and prints the station ID, temperature, dew point, and pressure.
func main() {
	stationID := "KSEA"

	// Get the observation details for KSEA
	obs, err := nws.GetLatestObservations(stationID)
	if err != nil {
		log.Fatalf("Failed to get observation station details: %v", err)
	}

	// Print the entire observationresponse object for debugging
	if debug {
		fmt.Printf("Observation Details: %+v\n\n", obs)
	}

	// Obtain the values
	temp := obs.Temperature.Value
	tempUnit := convertUnit(obs.Temperature.UnitCode)
	tempF := celsiusToFahrenheit(obs.Temperature.Value)
	dew := obs.Dewpoint.Value
	dewUnit := convertUnit(obs.Dewpoint.UnitCode)
	dewF := celsiusToFahrenheit(obs.Dewpoint.Value)
	pressure := obs.BarometricPressure.Value / 100

	// Print the values
	fmt.Printf("Station: %s\n", stationID)
	fmt.Printf("Temperature: %.1f%s - %.1f°F\n", temp, tempUnit, tempF)
	fmt.Printf("Dew Point: %.1f%s - %.1f°F\n", dew, dewUnit, dewF)
	fmt.Printf("Pressure: %.0f hPa\n", pressure)
}
