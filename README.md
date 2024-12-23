# nwsgo [![GoDoc](https://godoc.org/github.com/jacaudi/nws?status.svg)](https://godoc.org/github.com/jacaudi/nws)

## Purpose

Go package for the National Weather Service API. The data provided by weather.gov
is in the public domain and covers the continental United States. The service
is maintained by the National Weather Service under the umbrella of the
National Oceanic and Atmospheric Administration (NWS).

Data on various weather.gov API endpoints is measured at different intervals.
If a data point is measured hourly then you should take this into account when
polling for updates.

## API

`nws` is a Go client for the weather.gov API and supports the following endpoints:

```go
nwsgo.GetPoints(latlon string) // Get NWS Data at that point
nwsgo.RadarStations() // All Radar Sites -- Station Details
nwsgo.RadarStation(stationID string) // Radar Station Details
```


## Credits

This repo was forked from [icodealot/noaa](https://github.com/icodealot/noaa) and was used as a base for this repo. However, I decided to rebuild the library from scratch
