package functions

import "chatty/main/structs"

func HandleBroadcast(broadcast chan string, clients map[structs.Client]bool) {
	for {
		msg := <-broadcast
		for client := range clients {
			client.Messages <- msg
		}
	}
}
