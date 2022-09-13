package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	CONNECT_HOST = "localhost"
	CONNECT_PORT = "8080"
)

var GetRequestHandler = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World!"))
	},
)

var PostRequestHandler = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("It's a Post Request!"))
	},
)

var PathVariableHandler = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		name := vars["name"]

		//Determining request type
		if r.Method == http.MethodPut {
			w.Write([]byte("Hi " + name + ", this is a " + r.Method + " request"))
		} else {
			w.Write([]byte("Hi " + name + ", this is a " + r.Method + " request"))
		}
	},
)

func main() {
	router := mux.NewRouter()

	router.Handle("/", GetRequestHandler).Methods("GET")
	router.Handle("/post", PostRequestHandler).Methods("POST")
	//using same handler function for GET and PUT request.
	router.Handle("/hello/{name}", PathVariableHandler).Methods("GET", "PUT")

	fmt.Println("Starting Server at port :", CONNECT_PORT)
	log.Fatal(http.ListenAndServe(CONNECT_HOST+":"+CONNECT_PORT, router))
}
