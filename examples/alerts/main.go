package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/jacaudi/nwsgo"
)

var (
	debug, _ = strconv.ParseBool(os.Getenv("DEBUG"))
)

func main() {

	activeAlerts, err := nwsgo.GetActiveAlerts()
	if err != nil {
		log.Fatalf("Failed to get radar station details: %v", err)
	}

	// Print the entire activeAlerts object for debugging
	if debug {
		fmt.Printf("Active Alert Details: %+v\n\n", activeAlerts)
	}

	firstAlertDesc := activeAlerts.Data[0].Description

	// Print the Description of the First Alert Returned
	fmt.Printf("***-----ALERT EXAMPLE-----***\n%s\n***END***\n", firstAlertDesc)
}
