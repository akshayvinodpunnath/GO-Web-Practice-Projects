package main

import (
	"33-postgres-CRUD/pkg/routes"

	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	CONNECT_HOST = "localhost"
	CONNECT_PORT = "8080"
)

func main() {
	router := mux.NewRouter()
	routes.RegisterRoutes(router)

	fmt.Println("Starting server on port:", CONNECT_PORT)
	log.Fatal(http.ListenAndServe(CONNECT_HOST+":"+CONNECT_PORT, router))
}
