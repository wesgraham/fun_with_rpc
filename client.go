package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Clients have a ticker and a base. Role of Server is to receive client, and incriment ticker
type Client struct {
	Ticker int  `json:"ticker"`
	Base http.Client `json:"client"`
}

// Get current ticker value
func (client Client) getTicker() int {
	return client.Ticker
}

// Send request to server to increment ticker. Should return a response with incrimented ticker value
func (client Client) incrementTicker() int {
	// Send Current Client State
	body, err := json.Marshal(client.Ticker)
	if err != nil {
		fmt.Println("Error marshalling json")
		return -1
	}

	// Post Request
	resp, err := client.Base.Post("http://10.105.145.102:8080/", "application/json", bytes.NewBuffer(body))
	if err != nil {
		panic(err.Error())
		return -1
	} else if resp.StatusCode > 299 {
		fmt.Println("Bad status code")
		return -1
	}

	// Interpret Response
	var respbody int
	respbodybytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body")
		return -1
	}
	err = json.Unmarshal(respbodybytes, &respbody)
	if err != nil {
		fmt.Println("Error unmarshalling json")
		return -1
	}

	client.Ticker = respbody

	return client.Ticker
}

func (client Client) resetTicker() int {
	// Send Current Client State
	body, err := json.Marshal(client.Ticker)
	if err != nil {
		fmt.Println("Error marshalling json")
		return -1
	}

	// Post Request
	resp, err := client.Base.Post("http://10.105.145.102:8080/reset", "application/json", bytes.NewBuffer(body))
	if err != nil {
		panic(err.Error())
		return -1
	} else if resp.StatusCode > 299 {
		fmt.Println("Bad response code")
		return -1
	}

	// Interpret Response
	var respbody int
	respbodybytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body")
		return -1
	}
	err = json.Unmarshal(respbodybytes, &respbody)
	if err != nil {
		fmt.Println("Error unmarshalling response body")
		return -1
	}

	client.Ticker = respbody

	return client.Ticker
}