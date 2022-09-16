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

type Route struct {
	Name       string
	Method     string
	Path       string
	HandleFunc http.HandlerFunc
}

type Employee struct {
	Id        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type Routes []Route
type Employees []Employee

var routes Routes
var employees Employees

func init() {
	routes = Routes{
		Route{
			Name:       "getEmployees",
			Method:     "GET",
			Path:       "/employees",
			HandleFunc: getEmployees,
		},
		Route{
			Name:       "getEmployee",
			Method:     "GET",
			Path:       "/employees/{id}",
			HandleFunc: getEmployee,
		},
		Route{
			Name:       "createEmployee",
			Method:     "POST",
			Path:       "/employees",
			HandleFunc: createEmployee,
		},
		Route{
			Name:       "updateEmployee",
			Method:     "PUT",
			Path:       "/employees",
			HandleFunc: updateEmployee,
		},
		Route{
			Name:       "deleteEmployee",
			Method:     "DELETE",
			Path:       "/employees/{id}",
			HandleFunc: deleteEmployee,
		},
	}

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
	json, _ := json.Marshal(employees)
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

func getEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]
	for _, item := range employees {
		if item.Id == id {
			json, _ := json.Marshal(item)
			w.WriteHeader(http.StatusOK)
			w.Write(json)
			return
		}
	}
	json.NewEncoder(w).Encode("Id not found")
}

func createEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	employee := Employee{}
	err := json.NewDecoder(r.Body).Decode(&employee)
	if err != nil {
		log.Print("error occurred while decoding employee data :: ", err)
		return
	}
	employees = append(employees, employee)
	json, _ := json.Marshal(employee)
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

func updateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	employee := Employee{}
	err := json.NewDecoder(r.Body).Decode(&employee)
	if err != nil {
		log.Print("error occurred while decoding employee data :: ", err)
		return
	}

	for idx, item := range employees {
		if item.Id == employee.Id {
			employees[idx].FirstName = employee.FirstName
			employees[idx].LastName = employee.LastName
			json, _ := json.Marshal(employee)
			w.WriteHeader(http.StatusOK)
			w.Write(json)
			return
		}
	}
	json.NewEncoder(w).Encode("Id not found")

}

func deleteEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id := params["id"]
	for idx, item := range employees {
		if item.Id == id {
			employees = append(employees[:idx], employees[idx+1:]...)
			json, _ := json.Marshal(item)
			w.WriteHeader(http.StatusOK)
			w.Write(json)
			return
		}
	}
	json.NewEncoder(w).Encode("Id not found")
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
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./assets/")))

	fmt.Println("Starting server on port:", CONNECT_PORT)
	log.Fatal(http.ListenAndServe(CONNECT_HOST+":"+CONNECT_PORT, router))
}
