package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	CONNECT_HOST = "localhost"
	CONNECT_PORT = "8080"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func main() {
	http.HandleFunc("/", helloWorld)
	err := http.ListenAndServe(CONNECT_HOST+":"+CONNECT_PORT, nil)
	if err != nil {
		log.Fatal("error starting http server: ", err)
		return
	}
}
