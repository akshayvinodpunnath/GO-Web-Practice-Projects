package controllers

import (
	"33-postgres-CRUD/pkg/models"
	"strconv"

	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetEmployees(w http.ResponseWriter, r *http.Request) {
	data := models.GetEmployees()
	response := models.Response{}

	response.Data = data
	response.Message = "Employee Records"
	response.Status = true

	json, _ := json.Marshal(response)
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

func GetEmployee(w http.ResponseWriter, r *http.Request) {
	response := models.Response{}
	params := mux.Vars(r)
	id := params["id"]
	idVal, err := strconv.Atoi(id)
	if err != nil {
		response.Data = []models.Employee{}
		response.Message = err.Error()
		response.Status = false
		json, _ := json.Marshal(response)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(json)
		return
	}

	data := models.GetEmployee(idVal)

	response.Data = []models.Employee{data}
	response.Message = "Employee Record"
	response.Status = true

	json, _ := json.Marshal(response)
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

func CreateOrUpdateEmployee(w http.ResponseWriter, r *http.Request) {
	employee := models.Employee{}
	response := models.Response{}
	json.NewDecoder(r.Body).Decode(&employee)

	response.Data = []models.Employee{employee}

	if r.Method == http.MethodPost {
		err := models.CreateEmployee(employee)
		if err != nil {
			response.Message = err.Error()
			response.Status = false
			w.WriteHeader(http.StatusInternalServerError)

		} else {
			response.Message = "Employee Created"
			response.Status = true
			w.WriteHeader(http.StatusOK)
		}

	} else if r.Method == http.MethodPut {
		err := models.UpdateEmployee(employee)
		if err != nil {
			response.Message = err.Error()
			response.Status = false
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			response.Message = "Employee Updated"
			response.Status = true
			w.WriteHeader(http.StatusOK)
		}

	}

	json, _ := json.Marshal(response)
	w.Write(json)
}

func DeleteEmploye(w http.ResponseWriter, r *http.Request) {
	response := models.Response{}
	params := mux.Vars(r)
	id := params["id"]
	idVal, err := strconv.Atoi(id)
	if err != nil {
		response.Data = []models.Employee{}
		response.Message = err.Error()
		response.Status = false
		json, _ := json.Marshal(response)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(json)
		return
	}

	err = models.DeleteEmploye(idVal)
	if err != nil {
		response.Data = []models.Employee{}
		response.Message = err.Error()
		response.Status = false
		json, _ := json.Marshal(response)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(json)
	} else {
		response.Data = []models.Employee{
			{
				Id: idVal,
			},
		}
		response.Message = "Employee Deleted"
		response.Status = true
		json, _ := json.Marshal(response)
		w.WriteHeader(http.StatusOK)
		w.Write(json)
	}
}
