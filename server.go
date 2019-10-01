package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type InvalidRequest struct {
	message string
}

// Define Handler
func HandleIncrement(w http.ResponseWriter, req *http.Request) {
	// Decode request
	if req.Method != "POST" {
		respjson, err := json.Marshal(InvalidRequest{"must use POST"})
		if err != nil {
			// return error
		}
		// Respond with error
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		_, err = w.Write(respjson)
		if err != nil {
			fmt.Println("Error writing")
		}
	}

	var clientTicker int
	err := json.NewDecoder(req.Body).Decode(&clientTicker)
	if err != nil {
		// Respond with error
	}

	// Increment Ticker
	clientTicker += 1
	fmt.Printf("ticker value: %d \n", clientTicker)

	// Prepare response
	responsejson, err := json.Marshal(clientTicker)
	if err != nil {
		fmt.Println("Error marshalling json")
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(responsejson)
	if err != nil {
		fmt.Println("Error writing json")
	}
}

// Define Handler
func HandleReset(w http.ResponseWriter, req *http.Request) {
	// Decode request
	if req.Method != "POST" {
		respjson, err := json.Marshal(InvalidRequest{"must use POST"})
		if err != nil {
			// return error
		}
		// Respond with error
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		_, err = w.Write(respjson)
		if err != nil {
			// return error
		}
	}

	var client Client
	err := json.NewDecoder(req.Body).Decode(&client)
	if err != nil {
		// Respond with error
	}

	// Increment Ticker
	client.Ticker = 0

	// Prepare response
	responsejson, err := json.Marshal(client)
	if err != nil {
		// return error
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(responsejson)
	if err != nil {
		// return error
	}
}

// Setup Routes
func SetupRoutes() {
	http.HandleFunc("/", HandleIncrement)
	http.HandleFunc("/reset", HandleReset)
}
