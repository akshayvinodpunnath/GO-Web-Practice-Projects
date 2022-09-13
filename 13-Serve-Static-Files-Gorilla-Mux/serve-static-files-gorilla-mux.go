package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

const (
	CONNECT_HOST = "localhost"
	CONNECT_PORT = "8080"
)

type Person struct {
	Id   int
	Name string
}

func renderTemplate(w http.ResponseWriter, r *http.Request) {
	person := Person{Id: 44, Name: "Hamilton"}
	parseTemplate, _ := template.ParseFiles("templates/first-template.html")
	err := parseTemplate.Execute(w, person)
	if err != nil {
		log.Fatal("Error in parse template execution :", err.Error())
		return
	}
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", renderTemplate).Methods("GET")

	router.PathPrefix("/").Handler(http.StripPrefix("/static", http.FileServer(http.Dir("static/"))))

	fmt.Println("Starting server in port :", CONNECT_PORT)
	log.Fatal(http.ListenAndServe(CONNECT_HOST+":"+CONNECT_PORT, router))
}
