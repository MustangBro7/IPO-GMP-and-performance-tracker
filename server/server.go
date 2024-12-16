package main

import (
	"encoding/json"
	"fmt"
	ipo_tracker "ipo_tracker/ipos"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	headers, rows := ipo_tracker.GetGMP("https://www.investorgain.com/report/ipo-performance-history/486/ipo/", []int{0, 5, 6, 8}, "main")
	fmt.Println(headers)
	fmt.Println(rows)

	w.Header().Set("Content-Type", "application/json")

	// Encode the array of arrays into JSON
	if err := json.NewEncoder(w).Encode(rows); err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
	}
}

func main() {
	// Set up the route
	http.HandleFunc("/data", handler)

	// Start the server
	port := 8080
	fmt.Printf("Server is running on http://localhost:%d\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
