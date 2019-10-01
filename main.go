package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	mode := os.Args[1]
	fmt.Printf("You are running as a %s! \n", mode)
	if mode == "client" {
		if len(os.Args) < 2 {
			fmt.Println("Need to specify more arguments")
			os.Exit(1)
		}
		tickerVal, _ := strconv.Atoi(os.Args[2])
		fmt.Printf("Current ticker value = %d \n", tickerVal)
		client := Client{tickerVal, http.Client{Timeout:5*time.Second} }
		if len(os.Args) > 2 {
			for i := 3; i < len(os.Args); i++ {
				if os.Args[i] == "increment" {
					client.incrementTicker()
					fmt.Printf("Current ticker value = %d \n", client.getTicker())
				} else if os.Args[i] == "reset" {
					client.resetTicker()
					fmt.Printf("Current ticker value = %d \n", client.getTicker())
				}
			}
		}

	} else if mode == "server" {
		SetupRoutes()
		log.Fatal(http.ListenAndServe(":8080", nil))
	} else {
		fmt.Println("Invalid type specified")
	}
}