package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const (
	CONNECT_HOST = "localhost"
	CONNECT_PORT = "8080"
)

type Movie struct {
	Name     string  `json:"name"`
	Rating   float32 `json:"rating"`
	Director string  `json:"director"`
}

var movies []Movie

func helloWorld(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := "Hello World"
	json.NewEncoder(w).Encode(response)
}

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	res, _ := json.Marshal(movies)
	w.Write(res)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	name := r.URL.Query().Get("name")
	if name == "" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		json.NewEncoder(w).Encode("Invalid Params passed")
		return
	}

	w.WriteHeader(http.StatusOK)

	for _, item := range movies {
		if item.Name == name {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	json.NewEncoder(w).Encode("Name not found")
}

func main() {
	movies = append(movies, Movie{Name: "Flight", Rating: 7.5, Director: "Robert Zemeckis"})
	movies = append(movies, Movie{Name: "Gladiator", Rating: 8.5, Director: "Ridley Scott"})

	http.HandleFunc("/", helloWorld)
	http.HandleFunc("/movies", getMovies)
	http.HandleFunc("/movie", getMovie)
	fmt.Printf("Starting Server at Port: %s \n", CONNECT_PORT)
	log.Fatal(http.ListenAndServe(CONNECT_HOST+":"+CONNECT_PORT, nil))
}
