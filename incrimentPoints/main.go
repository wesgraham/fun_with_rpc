package incPoints

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"net/rpc/jsonrpc"
	"net/rpc"
	"net"
)

func main() {
	mode := os.Args[1]
	fmt.Printf("You are running as a %s! \n", mode)
	if mode == "client" {
		// Connect to server
		client, err := net.Dial("tcp", "127.0.0.1:8080")
		if err != nil {
			log.Fatal("dialing:", err)
		}

		// Setup Synchronous call
		price, err := strconv.Atoi(os.Args[3])
		item := &Item{3, price}
		var reply int
		rpcClient := jsonrpc.NewClient(client)
		// Call AddFive w/ params item and reply
		err = rpcClient.Call(os.Args[2], item, &reply)
		if err != nil {
			panic(err.Error())
		}
		item.Price = reply
		fmt.Printf("new item price = %d \n", item.Price)

	} else if mode == "server" {
		computer := new(Computer)
		server := rpc.NewServer()
		_ = server.Register(computer)
		server.HandleHTTP(rpc.DefaultRPCPath, rpc.DefaultDebugPath)
		listener, e := net.Listen("tcp", ":8080")
		if e != nil {
			log.Fatal("listen error:", e)
		}

		for {
			conn, _ := listener.Accept()
			log.Printf("new connection established\n")
			go server.ServeCodec(jsonrpc.NewServerCodec(conn))
		}

	} else {
		fmt.Println("Invalid type specified")
	}
}