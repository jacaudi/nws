package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func callAPI() {
	reader := bufio.NewReader(os.Stdin)

	// Prompt for URL
	fmt.Print("Enter the NWS API URL: ")
	url, _ := reader.ReadString('\n')
	url = strings.TrimSpace(url)

	// List of headers
	headers := []string{"application/ld+json", "application/geo+json"}

	// Display the list of headers
	fmt.Println("Select an Accept header from the list:")
	for i, header := range headers {
		fmt.Printf("%d: %s\n", i+1, header)
	}

	// Read user input for header selection
	fmt.Print("Enter the number of the header you want to use: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	// Convert input to an index
	index, err := strconv.Atoi(input)
	if err != nil || index < 1 || index > len(headers) {
		log.Fatalf("Invalid selection")
	}

	// Get the selected header
	selectedHeader := headers[index-1]

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	req.Header.Add("Accept", selectedHeader)

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error making request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v", err)
	}

	fmt.Println(string(body))
}

func main() {
	callAPI()
}
