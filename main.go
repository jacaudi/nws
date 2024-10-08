package main

import (
    "fmt"
    "github.com/jacaudi/nwsgo" // Adjust the import path according to your module
)

func main() {
    // Example function call
    stations, err := nws.RadarStationList()
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    fmt.Printf("Radar Stations: %+v\n", stations)
}

