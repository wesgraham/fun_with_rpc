package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	argv := os.Args

	if len(argv) < 2 {
		log.Fatal("insufficient arguments passed")
	}

	if argv[1] == "client" {

		// Create reply object
		var addReply int
		var getReply Potato
		var table map[string]Potato

		// Establish Connection
		connection, err := rpc.DialHTTP("tcp", ":8080")
		if err != nil {
			log.Fatal("Could not dial connection")
		}

		// Lets make some potatoes
		Irish := Potato{"Irish", "Irish", "Green", 21}
		Canadian := Potato{"Canadian", "Canadian", "Brown", 101}
		Old := Potato{"Old", "Old", "Grey", 1000}
		IrishNew := Potato{"Irish", "Irish", "Orange", 21}

		// Calls
		err = connection.Call("API.AddPotato", Irish, &addReply)
		err = connection.Call("API.AddPotato", Canadian, &addReply)
		err = connection.Call("API.AddPotato", Old, &addReply)
		err = connection.Call("API.GetPotato", Irish.Key, &getReply)
		fmt.Printf("Got: %s\n", getReply.Key);
		err = connection.Call("API.EditPotato", IrishNew, &addReply)
		fmt.Printf("Edited: %s\n", IrishNew.Key);
		err = connection.Call("API.DeletePotato", Old.Key, &addReply)
		fmt.Printf("Deleted: %s\n", Old.Key);
		err = connection.Call("API.DisplayPotatoTable", "", &table)
		fmt.Println(table)
		err = connection.Call("API.AddPotato", Irish, &addReply)
		if err != nil {
			fmt.Println("cant add potato that is already there")
		}

	} else if argv[1] == "server" {
		// Instance of api, register
		var api = new(API)
		err := rpc.Register(api)

		if err != nil {
			log.Fatal("Could not register api methods")
		}
		rpc.HandleHTTP()

		// Load latest value into PTable
		err = Load("./file.tmp", &PTable)
		if err != nil {
			log.Fatal(err.Error())
		}

		// Handle Signals Gracefully
		sigs := make(chan os.Signal, 1)
		done := make(chan bool, 1)

		signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)


		go func() {
			sig := <-sigs
			// Handle Signals
			fmt.Printf("handling, %s\n", sig)
			HandleSignal(sig)
			done <- true
		}()

		// Listen for connections
		listener, err := net.Listen("tcp", ":8080")
		if err != nil {
			log.Fatal("Could not listen")
		}
		fmt.Println("Successfully Spun Server");
		err = http.Serve(listener, nil)
		if err != nil {
			log.Fatal("Could not serve")
		}
	}
}
