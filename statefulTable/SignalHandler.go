package main

import (
	"fmt"
	"log"
	"os"
	"syscall"
)

func HandleSignal(sig os.Signal) {
	if sig == syscall.SIGINT {
		fmt.Println("sigint detected, saving state to file")

		// Save to file
		err := PTable.Save("./file.tmp")

		if err != nil {
			log.Fatal("Error saving table to state")
		}

		os.Exit(0)
	} else if sig == syscall.SIGTERM {
		fmt.Println("sigterm detected, saving state to file")

		// Save to file
		err := PTable.Save("./file.tmp")

		if err != nil {
			log.Fatal("Error saving table to state")
		}

		os.Exit(0)
	}
}
