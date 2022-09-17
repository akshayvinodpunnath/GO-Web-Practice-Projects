package controllers

import (
	"32-postgres-create-record/pkg/models"

	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func GetEmployees(w http.ResponseWriter, r *http.Request) {
	response := models.Response{}
	employees := models.GetEmployees()

	response.Data = employees
	response.Message = "Employee Records"
	response.Status = true
	json, _ := json.Marshal(response)
	w.WriteHeader(http.StatusOK)
	w.Write(json)

}

func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	employee := models.Employee{}
	response := models.Response{}
	err := json.NewDecoder(r.Body).Decode(&employee)
	if err != nil {
		fmt.Println("Error in decode")
		log.Fatal(err)
	}

	createEmployee := models.CreateEmployee(employee)

	if createEmployee {
		response.Data = []models.Employee{employee}
		response.Message = "Employee Created"
		response.Status = true
		json.NewEncoder(w).Encode(response)
	} else {
		response.Data = []models.Employee{}
		response.Message = "Id already exists"
		response.Status = true
		json.NewEncoder(w).Encode(response)
	}

}
