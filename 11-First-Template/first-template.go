package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
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
	person := Person{Id: 1, Name: "Max"}
	/*
		parsedTemplate, _ := template.ParseFiles("templates/first-template.html")
		Here we are calling ParseFiles of the html/template package, It creates a new template and parses the file passed as input.
		The resulting template will have the name and contents of the input file.
	*/
	parsedTemplate, _ := template.ParseFiles("templates/first-template.html")

	/*
		Execute handler on parsed template, which injects person data to template.
		Generates an HTML output and writes to HTTP response stream.
	*/
	err := parsedTemplate.Execute(w, person)
	if err != nil {
		log.Printf("Error occurred while executing the template or writing its output : ", err.Error())
		return
	}
}

func main() {
	http.HandleFunc("/", renderTemplate)

	fmt.Println("Starting Server at :", CONNECT_PORT)
	log.Fatal(http.ListenAndServe(CONNECT_HOST+":"+CONNECT_PORT, nil))
}
