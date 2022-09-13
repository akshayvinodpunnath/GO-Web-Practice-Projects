package main

import (
	"bufio"
	"fmt"
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
		log.Fatal("Error in starting TCP Server :", err.Error())
	}
	defer listener.Close()
	log.Println("Listing on " + CONNECT_HOST + " : " + CONNECT_PORT)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Error in listening", err.Error())
		}
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	message, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("Error reading: ", err.Error())
	}
	fmt.Println("Message received :", string(message))

	//conn.Write to send data back to client
	conn.Write([]byte(message + "\n"))
	conn.Close()
}
