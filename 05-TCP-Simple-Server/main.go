package main

import (
	"log"
	"net"
)

const (
	CONNECT_HOST = "localhost"
	CONNECT_PORT = "8080"
	CONNECT_TYPE = "tcp"
)

func main() {
	listener, err := net.Listen(CONNECT_TYPE, CONNECT_HOST+":"+CONNECT_PORT)
	if err != nil {
		log.Fatal("Server not starting :", err)
	}

	//defer statement closes a TCP socket listener when the application closes
	defer listener.Close()
	log.Println("Listening on " + CONNECT_HOST + ":" + CONNECT_PORT)

	/*
		Next, we accept the incoming request to the TCP server in a constant loop, and if there
		are any errors in accepting the request, then we log it and exit; otherwise, we simply
		print the connection object on the server console
	*/
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Error accepting: ", err.Error())
		}
		log.Println(conn)
	}
}
