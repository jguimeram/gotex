package main

import (
	"bufio"
	"log"
	"net"
)

type Client struct {
	name string
	conn net.Conn
	cmd  string
}

func (c *Client) help() {
	log.Println("from help function")
}

var messages = make(chan Client)

/* func listClients() {
	for _, c := range Clients {
		log.Println(c.name)
	}
} */

func NewClient(name string, conn net.Conn) Client {
	return Client{
		name: name,
		conn: conn,
	}
}

func startServer() {
	ln, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Println(err)
	}
	defer ln.Close()

	clients := make(map[string]Client)

	go broadcaster(clients)
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
		}
		go handleConnection(clients, conn)
	}
}

func handleConnection(clients map[string]Client, conn net.Conn) {
	defer conn.Close()

	//Clients := make(map[string]Client)

	//TODO: username issue
	c := NewClient("anonymous", conn)

	clients[conn.RemoteAddr().String()] = c

	input := bufio.NewScanner(conn)

	for input.Scan() {
		c.cmd = input.Text()
		//messages <- c
		switch c.cmd {
		case "/help":
			c.help()
		case "/quit":
			c.conn.Close()
		default:
			messages <- c
		}
	}

	delete(clients, conn.RemoteAddr().String())
}

func broadcaster(clients map[string]Client) {
	for {
		client := <-messages
		log.Println(clients)
		for _, c := range clients {
			if c.conn.RemoteAddr().String() == client.conn.RemoteAddr().String() {
				continue
			}
			c.conn.Write([]byte(client.cmd + "\n"))
		}

	}
}

func main() {
	startServer()

}
