package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/schema"
)

const (
	CONNECT_HOST = "localhost"
	CONNECT_PORT = "8080"
)

type User struct {
	Username string
	Password string
}

func readForm(r *http.Request) *User {
	r.ParseForm()
	user := new(User)
	decoder := schema.NewDecoder()
	decodeErr := decoder.Decode(user, r.PostForm)
	if decodeErr != nil {
		log.Printf("error mapping parsed form data to struct : ", decodeErr)
	}
	return user

}

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		parseFiles, _ := template.ParseFiles("templates/login-form.html")
		parseFiles.Execute(w, nil)
	} else {
		user := readForm(r)
		fmt.Fprintf(w, "Hello "+user.Username+"!")
	}

}

func main() {
	http.HandleFunc("/", login)

	fmt.Println("Starting server in port: ", CONNECT_PORT)
	log.Fatal(http.ListenAndServe(CONNECT_HOST+":"+CONNECT_PORT, nil))
}
