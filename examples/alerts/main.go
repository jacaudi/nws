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

	activeAlerts, err := nws.GetActiveAlerts()
	if err != nil {
		log.Fatalf("Failed to get alert details: %v", err)
	}

	// Print the entire activeAlerts object for debugging
	if debug {
		fmt.Printf("Active Alert Details: %+v\n\n", activeAlerts)
	}

	// Print the Description of the First Alert Returned
	fmt.Printf("***-----ALERT EXAMPLE-----***\n%s\n***END***\n", activeAlerts.Data[0].Description)

	// Pull the VTEC of the First 10 Alerts
	count := 0
	for _, alert := range activeAlerts.Data {
		if count >= 10 {
			break
		}
		if len(alert.Parameters.VTEC) == 0 {
			continue
		}
		fmt.Printf("VTEC of Alert %02d: %s\n", count+1, alert.Parameters.VTEC)
		count++
	}

	// Calculate the total number of alerts
	totalAlerts := len(activeAlerts.Data)
	fmt.Printf("Total number of alerts: %d\n", totalAlerts)
}
