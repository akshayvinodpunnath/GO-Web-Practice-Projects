package routes

import (
	"33-postgres-CRUD/pkg/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

var RegisterRoutes = func(router *mux.Router) {
	router.HandleFunc("/employees", controllers.GetEmployees).Methods(http.MethodGet)
	router.HandleFunc("/employees/{id}", controllers.GetEmployee).Methods(http.MethodGet)
	router.HandleFunc("/employees", controllers.CreateOrUpdateEmployee).Methods(http.MethodPost, http.MethodPut)
	router.HandleFunc("/employees/{id}", controllers.DeleteEmploye).Methods(http.MethodDelete)
}
