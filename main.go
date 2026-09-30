package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// User represents the data payload we will send back
type User struct {
	Name  string `json:"name"`
	Role  string `json:"role"`
}

func main() {
	// Create a new router (ServeMux)
	mux := http.NewServeMux()

	// 1. A simple text route
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Welcome to the Go backend server!")
	})

	// 2. A JSON response route
	mux.HandleFunc("GET /api/user", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		
		data := User{
			Name: "Alex",
			Role: "Developer",
		}

		// Encode the struct directly into the response writer
		json.NewEncoder(w).Encode(data)
	})

	// Define the port and start the server
	port := ":8080"
	fmt.Printf("Server starting smoothly on http://localhost%s\n", port)
	
	err := http.ListenAndServe(port, mux)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}