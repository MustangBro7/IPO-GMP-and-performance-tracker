package main

import (
	"encoding/json"
	"fmt"
	ipo_tracker "ipo_tracker/ipos"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {

	// Add CORS headers

	// Handle preflight OPTIONS request
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	// headers, rows := ipo_tracker.GetGMP("https://www.investorgain.com/report/ipo-performance-history/486/ipo/", []int{0, 5, 6, 8}, "main")
	headers, rows := ipo_tracker.Upcoming("https://www.investorgain.com/report/live-ipo-gmp/331/current/", []int{0, 1, 2, 3, 7, 8, 10})
	fmt.Println(headers)
	fmt.Println(rows)
	w.Header().Set("Access-Control-Allow-Origin", "*") // Allow requests only from your React app
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS") // Allow specific methods
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")       // Allow specific headers
	// Encode the array of arrays into JSON
	// if err := json.NewEncoder(w).Encode(headers); err != nil {
	// 	http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
	// }
	// data := map[string]interface{}{
	// 	"headers": headers,
	// 	"rows":    rows,
	// }
	if err := json.NewEncoder(w).Encode(rows); err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
	}

}

func main() {
	// Set up the route
	http.HandleFunc("/data", handler)

	// Start the server
	port := 8080
	fmt.Printf("Server is running on http://0.0.0.0:%d\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
