package main

import (
	"log"
	"net"
)

func startServer() {
	const PORT string = ":3000"

	ln, err := net.Listen("tcp", PORT)

	if err != nil {
		log.Panic("Could not start server. ", err)
	}
	defer ln.Close()

	//server welcome message
	log.Println("Connection to localhost chat succesfully")

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("Could not connect to server: %s", err)
			continue
		}
		go handleConnection(conn)
	}
}
