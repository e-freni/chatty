package main

import (
	"chatty/main/functions"
	"chatty/main/structs"
	"fmt"
	"net"
)

var Clients = make(map[structs.Client]bool)
var Broadcast = make(chan string)

func main() {

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()

	go functions.HandleBroadcast(Broadcast, Clients)

	fmt.Println("Chatty server started on :8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Connection error:", err)
			continue
		}
		go functions.HandleClient(conn, Clients, Broadcast)
	}
}
