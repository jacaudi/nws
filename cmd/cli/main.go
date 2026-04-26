package main

import (
	"fmt"
	"github.com/jacaudi/nws/cmd/nws"
)

func main() {
	alerts, err := nws.GetActiveAlerts()
	if err != nil {
		fmt.Println("Error fetching alerts:", err)
		return
	}
	fmt.Printf("Fetched %d active alerts. Updated: %s\n", len(alerts.Data), alerts.Updated)
	for _, alert := range alerts.Data {
		fmt.Println("---")
		fmt.Printf("ID: %s\nEvent: %s\nSeverity: %s\nHeadline: %s\n", alert.ID, alert.Event, alert.Severity, alert.Headline)
		fmt.Printf("  Sent: %s, Effective: %s, Expires: %s, Ends: %s\n", alert.Sent, alert.Effective, alert.Expires, alert.Ends)
		fmt.Println("  Description:", alert.Description)
		if alert.Instruction != "" {
			fmt.Println("  Instruction:", alert.Instruction)
		}
	}
}
