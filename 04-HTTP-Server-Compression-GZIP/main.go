package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
)

const (
	CONNECT_HOST = "localhost"
	CONNECT_PORT = "8080"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World\n")
}

func main() {
	mux := http.NewServeMux()

	//Endpoints
	mux.HandleFunc("/", helloWorld)

	fmt.Println("Starting server in port :", CONNECT_PORT)
	//compress the handler mux
	/*
		Here, we are calling http.ListenAndServe to serve HTTP requests that handle each
		incoming connection in a separate Goroutine for us. ListenAndServe accepts two
		parameters—server address and handler. Here, we are passing the server address
		as localhost:8080 and handler as CompressHandler, which wraps our server with a .gzip
		handler to compress all responses in a .gzip format.
	*/
	log.Fatal(http.ListenAndServe(CONNECT_HOST+":"+CONNECT_PORT, handlers.CompressHandler(mux)))
}
