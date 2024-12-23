package client

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// httpRequest makes an HTTP request to the NWS API and returns the response body.
func httpRequest(url string, agent string, accept string, units string, debug bool) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Accept", accept)

	if len(agent) == 0 {
		log.Panicf("The NWS API requires a User-Agent")
	} else {
		req.Header.Set("User-Agent", agent)
	}

	if len(units) == 0 {
		req.Header.Set("Units", "") // Defaults to US units if unspecified
	} else {
		req.Header.Set("Units", units)
	}

	if debug {
		log.Printf("Making request to URL: %s", url)
		log.Printf("Header: %s", accept)
		log.Printf("Header: %s", agent)
		log.Printf("Header: %s", units)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if debug {
		log.Printf("Received response status: %s", resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if debug {
		log.Printf("Response body: %s", string(body))
	}

	return body, nil
}
