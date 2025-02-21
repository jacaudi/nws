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

	wfo := pointData.GridID
	gridpoint := fmt.Sprintf("%s,%s", strconv.FormatFloat(pointData.GridX, 'f', -1, 64), strconv.FormatFloat(pointData.GridY, 'f', -1, 64))

	if debug {
		log.Printf("wfo: %s", wfo)
		log.Printf("gridpoint: %s", gridpoint)
	}

	// Get the radar station details for KATX
	forecastResponse, err := nws.GetForecast(wfo, gridpoint)
	if err != nil {
		log.Fatalf("Failed to get radar station details: %v", err)
	}

	if debug {
		log.Printf("Forecast Endpoint Response: %v\n\n", forecastResponse)
	}

	todayForecast := fmt.Sprintf("%s -- %s\n", forecastResponse.Periods[0].Name, forecastResponse.Periods[0].ShortForecast)
	allForecast := ""
	for _, period := range forecastResponse.Periods {
		if period.IsDaytime {
			allForecast += fmt.Sprintf("%s -- High Temperature: %s\n", period.Name, strconv.Itoa(period.Temperature))
		} else {
			allForecast += fmt.Sprintf("%s -- Low Temperature: %s\n", period.Name, strconv.Itoa(period.Temperature))
		}
	}

	// Print the Short Forecast
	fmt.Printf("Forecast: %s\n", todayForecast)
	fmt.Printf("All Forecasts:\n%s", allForecast)
}
