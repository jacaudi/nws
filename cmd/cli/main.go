package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jacaudi/nws/cmd/nws"
)

func main() {
	client, err := nws.NewClient(
		nws.WithUserAgent("nws-cli/1.0 (+https://github.com/jacaudi/nws)"),
	)
	if err != nil {
		log.Fatalf("NewClient: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	active, err := client.GetActiveAlerts(ctx)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error fetching alerts:", err)
		os.Exit(1)
	}

	fmt.Printf("Fetched %d active alerts. Updated: %s\n", len(active.Data), active.Updated)
	for _, alert := range active.Data {
		fmt.Println("---")
		fmt.Printf("ID: %s\nEvent: %s\nSeverity: %s\nHeadline: %s\n", alert.ID, alert.Event, alert.Severity, alert.Headline)
		fmt.Printf("  Sent: %s, Effective: %s, Expires: %s, Ends: %s\n", alert.Sent, alert.Effective, alert.Expires, alert.Ends)
		fmt.Println("  Description:", alert.Description)
		if alert.Instruction != "" {
			fmt.Println("  Instruction:", alert.Instruction)
		}
	}
}
