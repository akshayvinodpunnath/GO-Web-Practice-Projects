package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
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

var ExtraPostRequestHandler = http.HandlerFunc(
	func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("It's a Extra Post Request!"))
	},
)

func main() {
	router := mux.NewRouter()

	//os.Stdout - passed a standard output stream as a writer to it,
	//which means we are simply asking to log every request with the URL path / on the console in Apache Common Log Format
	router.Handle("/", handlers.LoggingHandler(os.Stdout, http.HandlerFunc(GetRequestHandler))).Methods("GET")

	//create a logfile server.log
	logFile, err := os.OpenFile("server.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("error starting http server : ", err.Error())
		return
	}

	//LoggingHandler for logging HTTP requests in the Apache Common Log Format.
	router.Handle("/post", handlers.LoggingHandler(logFile, PostRequestHandler)).Methods("POST")

	//CombinedLoggingHandler for logging HTTP requests in the Apache Combined Log Format commonly used by both Apache and nginx.
	router.Handle("/extraPost", handlers.CombinedLoggingHandler(logFile, ExtraPostRequestHandler)).Methods("POST")
	router.Handle("/hello/{name}", handlers.CombinedLoggingHandler(logFile, PathVariableHandler)).Methods("GET", "PUT")

	fmt.Println("Starting Server at port :", CONNECT_PORT)
	log.Fatal(http.ListenAndServe(CONNECT_HOST+":"+CONNECT_PORT, router))
}
