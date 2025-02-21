# nws

A Go library package built around the National Weather Service API. The data provided by weather.gov is in the public domain and covers the United States. The service is maintained by the National Weather Service under the umbrella of the National Oceanic and Atmospheric Administration (NOAA).

Data on various weather.gov API endpoints is measured at different intervals.
If a data point is measured hourly then you should take this into account when
polling for updates.

## API

`nws` is a Go client for the weather.gov API and supports the following endpoints:

```go
nws.GetPoints(latlon string) // Get NWS Data at that point
nws.RadarStations() // All Radar Sites -- Station Details
nws.RadarStation(stationID string) // Radar Station Details
```

## Credits

This repo was forked from [icodealot/noaa](https://github.com/icodealot/noaa) and was used as a base for this repo. However, I decided to rebuild the library from scratch
