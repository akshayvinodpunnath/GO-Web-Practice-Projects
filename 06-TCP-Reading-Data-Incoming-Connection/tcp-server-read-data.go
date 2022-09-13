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
		log.Fatal("Unable to start server: ", err.Error())
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Error Accepting connections : ", err.Error())
		}
		//handleRequest from the main() method using the go keyword, which means we are invoking a function in a Goroutine.
		go handleRequest(conn)
	}
}

/*
	we defined the handleRequest function, which reads an incoming connection
	into the buffer until the first occurrence of \n and prints the message on the
	console. If there are any errors in reading the message then it prints the error
	message along with the error object and finally closes the connection
*/
func handleRequest(conn net.Conn) {
	message, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	fmt.Print("Message Received from the client: ", string(message))
	conn.Close()
}
