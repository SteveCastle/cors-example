package main

import (
	"fmt"
	"net/http"
)

// This example demonstrates how to use the http package to create a web server
// that renders a single endpoint, /api, which returns a JSON response
// containing the message "The request succeeded!".

func main () {
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8081")
		switch r.Method {
		case "OPTIONS":
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
			w.WriteHeader(http.StatusOK)
		case "GET":
			// Print the request body and indicate we are handling a GET request.
			fmt.Println("Handling GET request")
			fmt.Fprintf(w, `{"message": "The request succeeded!"}`)
		case "POST":
			// Print the request body and indicate we are handling a POST request.
			fmt.Println("Handling POST request")
			fmt.Fprintf(w, `{"message": "The request succeeded!"}`)
		default:
			http.Error(w, "Invalid request method.", http.StatusMethodNotAllowed)
		}
	})
	http.ListenAndServe(":8080", nil)
}

