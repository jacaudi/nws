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
		nws.WithUserAgent("nws-example-forecast/1.0 (+https://github.com/jacaudi/nws)"),
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

	wfo := pointData.GridID
	gridpoint := fmt.Sprintf("%s,%s",
		strconv.FormatFloat(pointData.GridX, 'f', -1, 64),
		strconv.FormatFloat(pointData.GridY, 'f', -1, 64),
	)
	if debug {
		log.Printf("wfo=%s gridpoint=%s", wfo, gridpoint)
	}

	forecast, err := client.GetForecast(ctx, wfo, gridpoint)
	if err != nil {
		log.Fatalf("GetForecast: %v", err)
	}
	if debug {
		log.Printf("Forecast: %+v", forecast)
	}

	today := fmt.Sprintf("%s -- %s\n", forecast.Periods[0].Name, forecast.Periods[0].ShortForecast)
	all := ""
	for _, p := range forecast.Periods {
		label := "Low Temperature"
		if p.IsDaytime {
			label = "High Temperature"
		}
		all += fmt.Sprintf("%s -- %s: %s\n", p.Name, label, strconv.Itoa(p.Temperature))
	}
	fmt.Printf("Forecast: %s\n", today)
	fmt.Printf("All Forecasts:\n%s", all)
}
