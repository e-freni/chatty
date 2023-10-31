package structs

import (
	"fmt"
	"net"
)

type Client struct {
	Conn     net.Conn
	Name     string
	Messages chan string
}

func (c Client) WriteMessages() {
	for msg := range c.Messages {
		fmt.Fprintln(c.Conn, msg)
	}
}
