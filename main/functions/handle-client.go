package functions

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"
)

import "chatty/main/structs"

func HandleClient(conn net.Conn, clients map[structs.Client]bool, broadcast chan string) {
	messages := make(chan string)

	fmt.Fprintln(conn, "Welcome to this chat server! Please insert your name:")

	scanner := bufio.NewScanner(conn)
	scanner.Scan()
	clientName := scanner.Text()

	client := structs.Client{Conn: conn, Name: clientName, Messages: messages}

	fmt.Fprintln(conn, "Great, now you can chat to anyone who's connected to localhost:8080 (use nc or telnet), have fun!")

	go client.WriteMessages()

	clients[client] = true
	defer delete(clients, client)

	for scanner.Scan() {
		msg := scanner.Text()
		if strings.TrimSpace(msg) == "/exit" {
			break
		}
		broadcast <- fmt.Sprintf("[%s] - %s: %s", client.Name, time.Now().Format("15:04:05"), msg)
	}
	close(client.Messages)
	conn.Close()
}
