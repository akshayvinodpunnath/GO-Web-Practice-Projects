package routes

import (
	"32-postgres-create-record/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegisterRoutes = func(router *mux.Router) {
	router.HandleFunc("/employees", controllers.GetEmployees).Methods("GET")
	router.HandleFunc("/employees", controllers.CreateEmployee).Methods("POST")
}
