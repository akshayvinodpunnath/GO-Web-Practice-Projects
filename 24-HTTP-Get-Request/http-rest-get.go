package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	CONNECT_HOST = "localhost"
	CONNECT_PORT = "8080"
)

//Route struct
type Route struct {
	Name       string
	Method     string
	Path       string
	HandleFunc http.HandlerFunc
}

type Routes []Route

//route list
var routes = Routes{
	Route{
		"getEmployes",
		"GET",
		"/employees",
		getEmployees,
	},
	Route{
		"getEmployee",
		"GET",
		"/employees/{id}",
		getEmployee,
	},
}

type Employee struct {
	Id        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type Employees []Employee

var employees Employees

func init() {
	employees = Employees{
		Employee{
			Id:        "1",
			FirstName: "Max",
			LastName:  "Vestappen",
		},
		Employee{
			Id:        "44",
			FirstName: "Lewis",
			LastName:  "Hamilton",
		},
	}
}

func getEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	res, _ := json.Marshal(employees)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func getEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	for _, item := range employees {
		if item.Id == id {
			res, _ := json.Marshal(item)
			w.WriteHeader(http.StatusOK)
			w.Write(res)
			return
		}
	}
	json.NewEncoder(w).Encode("Id not found")
}

func addRoutes(router *mux.Router) *mux.Router {
	for _, route := range routes {
		router.
			Methods(route.Method).
			Name(route.Name).
			Path(route.Path).
			Handler(route.HandleFunc)
	}
	return router
}

func main() {
	muxRouter := mux.NewRouter().StrictSlash(true)
	router := addRoutes(muxRouter)
	fmt.Println("Starting server in port:", CONNECT_PORT)
	log.Fatal(http.ListenAndServe(CONNECT_HOST+":"+CONNECT_PORT, router))
}
