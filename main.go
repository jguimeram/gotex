package main

import (
	"bufio"
	"fmt"
	"net"
)

var clients = make(map[string]Client)
var messages = make(chan string)

/* var messages = make(chan Message)
var quit = make(chan Message)

type Message struct {
	text    string
	address string
} */

type Client struct {
	name   string
	conn   net.Conn
	status bool
}

func handleConnection(conn net.Conn) {

	defer conn.Close()

	client := Client{conn: conn} //initiates client
	msg := []byte("Welcome to the chat\n")
	conn.Write(msg)

	input := bufio.NewScanner(conn)

	conn.Write([]byte("Enter a username:\n"))
	input.Scan()
	username := input.Text()
	client.name = username

	for _, client := range clients {
		if username == client.name {
			conn.Write([]byte("User already exists!\n"))
			return
		}
	}

	clients[conn.RemoteAddr().String()] = client

	for input.Scan() {

		messages <- input.Text()

	}
}

func broadcaster() {
	for {
		msg := <-messages
		fmt.Println(msg)
	}
}

func main() {
	go broadcaster()
	startServer()
}
