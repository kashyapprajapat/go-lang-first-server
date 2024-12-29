package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

// Handler for the root path
func handler(w http.ResponseWriter, r *http.Request) {
	// Respond with a simple message
	fmt.Fprintf(w, "Hello, World!")
}

func SumFrom1To10Million(w http.ResponseWriter, r *http.Request) {
	// Calculate the sum
	total := 0
	for i := 1; i <= 10000000; i++ {
		total += i
	}

	// Return the result to the client
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("The total sum from 1 to 10 million is: " + strconv.Itoa(total)))
}

func SumFrom1To10Crore(w http.ResponseWriter, r *http.Request) {
	// Calculate the sum
	total := 0
	for i := 1; i <= 100000000; i++ { // Looping until 10 crore (100 million)
		total += i
	}

	// Return the result to the client
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("The total sum from 1 to 10 crore is: " + strconv.Itoa(total)))
}

func main() {
	// Use the PORT environment variable or default to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		fmt.Println("PORT environment variable not set. Defaulting to :8080")
	}

	// Set up routes
	http.HandleFunc("/", handler)
	http.HandleFunc("/loop", SumFrom1To10Million)
	http.HandleFunc("/loop2", SumFrom1To10Crore)

	// Start the server
	fmt.Printf("Server starting on port %s...\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
