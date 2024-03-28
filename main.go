package main

import (
	"bufio"
	"net"
)

var clients = make(map[string]Client)
var messages = make(chan Message)

/* var messages = make(chan Message)
var quit = make(chan Message)

type Message struct {
	text    string
	address string
} */

type Client struct {
	name string
	conn net.Conn
}

type Message struct {
	text   string
	client string
}

func handleConnection(conn net.Conn) {

	defer conn.Close()

	client := Client{conn: conn} //initiates client
	msg := Message{}
	welcome := []byte("Welcome to the chat\n")
	conn.Write(welcome)

	input := bufio.NewScanner(conn)

	//TODO: ask for a user
	conn.Write([]byte("Enter a username:\n"))
	input.Scan()
	username := input.Text()
	client.name = username //TODO: Clients struct
	msg.client = username  //TODO: Messages struct

	//TODO: check if users exists
	for _, client := range clients {
		if username == client.name {
			conn.Write([]byte("User already exists!\n"))
			return
		}
	}

	//TODO: if does not exists, add to map
	clients[conn.RemoteAddr().String()] = client

	//TODO: sending messages to chat room
	for input.Scan() {
		msg.text = input.Text()
		messages <- msg
	}
}

func broadcaster() {
	for {
		msg := <-messages
		for _, client := range clients {
			if msg.client == client.name {
				continue
			}
			client.conn.Write([]byte(msg.text + "\n"))
		}
	}
}

func main() {
	go broadcaster()
	startServer()
}
