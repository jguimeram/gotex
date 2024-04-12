package main

import (
	"log"
	"net"
)

type Client struct {
	name string
	conn net.Conn
}

var Clients map[string]Client

func listClients() {
	for _, c := range Clients {
		log.Println(c.name)
	}
}

func NewClient(name string, conn net.Conn) *Client {
	return &Client{
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
	c := Client{name: "anonymous", conn: conn}

	Clients[conn.RemoteAddr().Network()] = c

	/* var msg string
	input := bufio.NewScanner(conn)
	for input.Scan() {
		msg = input.Text()
		log.Println(msg)
	} */
}

func main() {
	startServer()
}
