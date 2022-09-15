package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

const (
	CONNECT_HOST = "localhost"
	CONNECT_PORT = "8080"
)

type Route struct {
	Name       string
	Method     string
	Path       string
	HandleFunc http.HandlerFunc
}

type Employee struct {
	Id        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:lastName`
}

type Routes []Route
type Employees []Employee

var routes Routes
var employees Employees
var employeesV1 Employees
var employeesV2 Employees

func init() {
	routes = Routes{
		Route{
			Name:       "getEmployees",
			Method:     "GET",
			Path:       "/employees",
			HandleFunc: getEmployees,
		},
	}

	employees = Employees{
		Employee{
			Id:        "1",
			FirstName: "Max",
			LastName:  "Vestappen",
		},
	}

	employeesV1 = Employees{
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

	employeesV2 = Employees{
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
		Employee{
			Id:        "5",
			FirstName: "Carlos",
			LastName:  "Sainz",
		},
	}
}

func getEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if strings.HasPrefix(r.URL.Path, "/v1") {
		json.NewEncoder(w).Encode(employeesV1)
	} else if strings.HasPrefix(r.URL.Path, "/v2") {
		json.NewEncoder(w).Encode(employeesV2)
	} else {
		json.NewEncoder(w).Encode(employees)
	}
}

func addRoutes(router *mux.Router) *mux.Router {
	for _, route := range routes {
		router.
			Name(route.Name).
			Methods(route.Method).
			Path(route.Path).
			Handler(route.HandleFunc)
	}
	return router
}

func main() {
	muxRouter := mux.NewRouter().StrictSlash(true)
	router := addRoutes(muxRouter)
	addRoutes(muxRouter.PathPrefix("/v1").Subrouter())
	addRoutes(muxRouter.PathPrefix("/v2").Subrouter())
	fmt.Println("Starting server in port:", CONNECT_PORT)
	log.Fatal(http.ListenAndServe(CONNECT_HOST+":"+CONNECT_PORT, router))
}
