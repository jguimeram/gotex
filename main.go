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
	go broadcaster()
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	Clients := make(map[string]Client)

	//TODO: username issue
	c := NewClient("anonymous", conn)

	Clients[conn.RemoteAddr().String()] = c

	input := bufio.NewScanner(conn)

	for input.Scan() {
		c.cmd = input.Text()
		//messages <- c
		switch c.cmd {
		case "/help":
			c.help()
		default:
			messages <- c
		}
	}
}

func broadcaster() {
	for {
		client := <-messages
		log.Println(client.cmd)
	}
}

func main() {
	startServer()

}
