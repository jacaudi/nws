package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/jacaudi/nwsgo"
)

var (
	debug, _  = strconv.ParseBool(os.Getenv("DEBUG"))
	radarMode = ""
)

func main() {

	activeAlerts, err := nwsgo.GetActiveAlerts()
	if err != nil {
		log.Fatalf("Failed to get radar station details: %v", err)
	}

	// Print the entire radarStation object for debugging
	if debug {
		fmt.Printf("Active Alert Details: %+v\n\n", activeAlerts)
	}

	// Print the VolumeCoveragePattern
	fmt.Printf("ALERTS: %s\n", activeAlerts)
}
