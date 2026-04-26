# nws

A Go library for the U.S. National Weather Service API
([api.weather.gov](https://api.weather.gov)). Data from weather.gov is in
the public domain and covers the United States. The service is operated
by the National Weather Service, an agency of NOAA.

Data on different weather.gov endpoints is measured at different
intervals — for hourly data, do not poll faster than once an hour.

## Install

```
go get github.com/jacaudi/nws
```

Module path: `github.com/jacaudi/nws`. Import package:
`github.com/jacaudi/nws/cmd/nws`.

## Usage

### Idiomatic: `*Client` with context

```go
package main

import (
    "context"
    "log"
    "time"

    "github.com/jacaudi/nws/cmd/nws"
)

func main() {
    client, err := nws.NewClient(
        nws.WithUserAgent("myapp/1.0 (contact@example.com)"),
    )
    if err != nil {
        log.Fatal(err)
    }

    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()

    station, err := client.RadarStation(ctx, "KATX")
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("Radar: %s", station.Name)
}
```

The NWS API requires a `User-Agent` that identifies your application.

### Shorthand: package-level functions

For one-shot scripts where the default `Client` is fine:

```go
station, err := nws.RadarStation("KATX")
```

These call `nws.DefaultClient` with `context.Background()`. Mutate
`DefaultClient` once at program start if you need to override the
default User-Agent globally.

## Endpoints

| Method | Endpoint |
|---|---|
| `GetPoints(ctx, latlon)` | `/points/{lat,lon}` |
| `RadarStations(ctx)` | `/radar/stations` |
| `RadarStation(ctx, id)` | `/radar/stations/{id}` |
| `GetForecast(ctx, wfo, gridpoint)` | `/gridpoints/{wfo}/{gridpoint}/forecast` |
| `GetActiveAlerts(ctx)` | `/alerts/active` |
| `GetLatestObservations(ctx, id)` | `/stations/{id}/observations/latest` |

## Errors

Non-2xx responses are returned as `*nws.APIError`:

```go
station, err := client.RadarStation(ctx, "BADSTATION")

var apiErr *nws.APIError
if errors.As(err, &apiErr) {
    if apiErr.StatusCode == http.StatusNotFound {
        // handle missing station
    }
}
```

## Configuration

`NewClient` accepts functional options:

| Option | Purpose |
|---|---|
| `WithBaseURL(url)` | Override the API base URL (default: `https://api.weather.gov`) |
| `WithUserAgent(ua)` | Override the User-Agent header |
| `WithAccept(mt)` | Override the Accept header (default: `application/ld+json`) |
| `WithUnits(u)` | Set unit system: `"us"`, `"si"`, or `""` |
| `WithHTTPClient(*http.Client)` | Provide a custom HTTP client (custom transport, retries, etc.) |

Invalid options return an error. The default `*http.Client` has a
30-second timeout.

## Versioning

This library follows semver. v0.1.0 introduced the `*Client` API; the
package-level wrapper functions are preserved across the change.

## Credits

This repo was forked from [icodealot/noaa](https://github.com/icodealot/noaa)
as a starting point and rebuilt from scratch.

## License

See [LICENSE](./LICENSE).
