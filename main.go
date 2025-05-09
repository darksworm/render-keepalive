package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	// Get URL from environment variable
	url := os.Getenv("PING_URL")
	if url == "" {
		fmt.Println("Error: PING_URL environment variable is not set")
		os.Exit(1)
	}

	fmt.Printf("Starting to ping URL: %s\n", url)

	// Ping the URL once without waiting for a response
	pingURLAsync(url)

	fmt.Println("Request sent successfully")
}

func pingURLAsync(url string) {
	fmt.Printf("Sending request to %s without waiting for response...\n", url)

	// Start a goroutine to make the HTTP request asynchronously
	go func() {
		_, err := http.Get(url)
		if err != nil {
			// Just log the error, but don't exit the program as the main goroutine has already exited
			fmt.Printf("Error sending request to %s: %v\n", url, err)
		}
		// We don't care about the response, so we don't process it
	}()

	// Return immediately without waiting for the request to complete
}
