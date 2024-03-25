package main

import (
	"bufio"
	"log"
	"net"
)

var clients = make(map[string]net.Conn)
var messages = make(chan Message)
var quit = make(chan Message)

type Message struct {
	text    string
	address string
}

type Client struct {
	name string
	conn net.Conn
}

func newUser(name string, conn net.Conn) Client {
	return Client{
		name: name,
		conn: conn,
	}
}

func messageBuilder(msg string, conn net.Conn) Message {
	return Message{
		text:    conn.RemoteAddr().String() + msg,
		address: conn.RemoteAddr().String(),
	}
}

func clientConnection(conn net.Conn) {
	defer conn.Close()

	//server log
	log.Printf("User %s join the channel", conn.RemoteAddr().String())

	//mapping user connection
	clients[conn.RemoteAddr().String()] = conn

	messages <- messageBuilder(" join the channel\n", conn)

	input := bufio.NewScanner(conn)

	for input.Scan() {
		messages <- messageBuilder(": "+input.Text()+"\n", conn)
	}

	delete(clients, conn.RemoteAddr().String())

	quit <- messageBuilder(" leave the chat\n", conn)
	log.Printf("User %s leave the channel", conn.RemoteAddr().String())

}

func broadcast() {
	for {
		select {
		//msg get the struct from message channel
		case msg := <-messages:
			//print messages to other clients
			for _, client := range clients {
				if msg.address == client.RemoteAddr().String() {
					continue
				}
				_, err := client.Write([]byte(msg.text))
				if err != nil {
					log.Println("Message not sent succesfuly")
				}
			}
		case msg := <-quit:
			for _, client := range clients {
				_, err := client.Write([]byte(msg.text))
				if err != nil {
					log.Println("Message not sent succesfuly")
				}
			}
		}
	}
}

func main() {

	const PORT string = ":3000"

	ln, err := net.Listen("tcp", PORT)

	if err != nil {
		log.Panic("Could not connect to server. ", err)
	}
	defer ln.Close()

	//server welcome message
	log.Println("Connection to localhost chat succesfully")

	go broadcast()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("Could not connect to server: %s", err)
			continue
		}
		go clientConnection(conn)
	}
}
