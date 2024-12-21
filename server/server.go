package main

import (
	"encoding/json"
	"fmt"
	ipo_tracker "ipo_tracker/ipos"
	"net/http"
)

func upcominghandler(w http.ResponseWriter, r *http.Request) {

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
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")                       // Allow specific methods
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, ngrok-skip-browser-warning") // Allow specific headers

	data := map[string]interface{}{
		"headers": headers,
		"rows":    rows,
	}
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
	}

}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	headers, rows := ipo_tracker.GetGMP("https://www.investorgain.com/report/ipo-performance-history/486/ipo/", []int{0, 5, 6, 8}, "main")
	fmt.Println(headers)
	fmt.Println(rows)
	w.Header().Set("Access-Control-Allow-Origin", "*") // Allow requests only from your React app
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS") // Allow specific methods
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")       // Allow specific headers

	data := map[string]interface{}{
		"headers": headers,
		"rows":    rows,
	}
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
	}
}

func smeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	headers, rows := ipo_tracker.GetGMP("https://www.investorgain.com/report/ipo-performance-history/486/sme/", []int{0, 5, 6, 8}, "sme")
	fmt.Println(headers)
	fmt.Println(rows)
	w.Header().Set("Access-Control-Allow-Origin", "*") // Allow requests only from your React app
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS") // Allow specific methods
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")       // Allow specific headers

	data := map[string]interface{}{
		"headers": headers,
		"rows":    rows,
	}
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
	}
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, ngrok-skip-browser-warning")

		// Handle preflight requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	// Set up the route
	http.Handle("/data", corsMiddleware(http.HandlerFunc(upcominghandler)))
	http.Handle("/main", corsMiddleware(http.HandlerFunc(mainHandler)))
	http.Handle("/sme", corsMiddleware(http.HandlerFunc(smeHandler)))

	// Start the server
	port := 8080
	fmt.Printf("Server is running on http://0.0.0.0:%d\n", port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
