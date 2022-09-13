package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

const (
	CONNECT_HOST = "localhost"
	CONNECT_PORT = "8080"
)

func login(w http.ResponseWriter, r *http.Request) {
	parseTemplate, _ := template.ParseFiles("templates/login-form.html")
	parseTemplate.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", login)

	fmt.Println("Starting server at port: ", CONNECT_PORT)
	log.Fatal(http.ListenAndServe(CONNECT_HOST+":"+CONNECT_PORT, nil))
}
