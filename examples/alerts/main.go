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
		nws.WithUserAgent("nws-example-alerts/1.0 (+https://github.com/jacaudi/nws)"),
	)
	if err != nil {
		log.Fatalf("NewClient: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	active, err := client.GetActiveAlerts(ctx)
	if err != nil {
		log.Fatalf("GetActiveAlerts: %v", err)
	}
	if debug {
		fmt.Printf("Active Alert Details: %+v\n\n", active)
	}

	if len(active.Data) > 0 {
		fmt.Printf("***-----ALERT EXAMPLE-----***\n%s\n***END***\n", active.Data[0].Description)
	} else {
		fmt.Println("No active alerts.")
	}

	count := 0
	for _, a := range active.Data {
		if count >= 10 {
			break
		}
		if len(a.Parameters.VTEC) == 0 {
			continue
		}
		fmt.Printf("VTEC of Alert %02d: %s\n", count+1, a.Parameters.VTEC)
		count++
	}

	fmt.Printf("Total number of alerts: %d\n", len(active.Data))
}
