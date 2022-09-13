package main

//crypto/subtle will be used to compare user entered username and password against ADMIN_USER and ADMIN_PASSWORD

import (
	"crypto/subtle"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const (
	CONNECT_HOST   = "localhost"
	CONNECT_PORT   = "8080"
	ADMIN_USER     = "admin"
	ADMIN_PASSWORD = "admin"
)

type Movie struct {
	Name     string  `json:"name"`
	Rating   float32 `json:"rating"`
	Director string  `json:"director"`
}

var movies []Movie

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func BasicAuthenticator(handler http.HandlerFunc, realm string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//r.BasicAuth() returns username password entered
		user, pass, ok := r.BasicAuth()
		if !ok || subtle.ConstantTimeCompare([]byte(user), []byte(ADMIN_USER)) != 1 || subtle.ConstantTimeCompare([]byte(pass), []byte(ADMIN_PASSWORD)) != 1 {
			w.Header().Set("WWW-Authenticate", `Basic realm="`+realm+`"`)
			w.WriteHeader(401)
			w.Write([]byte("You are Unauthorized to access the application.\n"))
			return
		}
		handler(w, r)
	}
}

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	res, _ := json.Marshal(movies)
	w.Write(res)
}

func main() {
	movies = append(movies, Movie{Name: "Flight", Rating: 7.5, Director: "Robert Zemeckis"})
	movies = append(movies, Movie{Name: "Gladiator", Rating: 8.5, Director: "Ridley Scott"})

	http.HandleFunc("/", BasicAuthenticator(helloWorld, "Please enter your username and password"))
	http.HandleFunc("/movies", BasicAuthenticator(getMovies, "Please enter your username and password"))
	fmt.Println("Starting Server on Port: ", CONNECT_PORT)
	log.Fatal(http.ListenAndServe(CONNECT_HOST+":"+CONNECT_PORT, nil))
}
