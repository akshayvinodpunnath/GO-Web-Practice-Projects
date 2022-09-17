package models

import (
	"33-postgres-CRUD/pkg/config"
	"errors"
	"fmt"
	"log"

	"github.com/jackc/pgx"
)

type Employee struct {
	Id        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

type Response struct {
	Status  bool       `json:"status"`
	Message string     `json:"message"`
	Data    []Employee `json:"data"`
}

var db *pgx.Conn

func init() {
	config.Connect()
	db = config.GetDB()
	config.CreateTables()
}

func GetEmployees() []Employee {
	rows, err := db.Query("SELECT * from EMPLOYEES")
	if err != nil {
		log.Fatal(err)
	}

	var employees = []Employee{}
	for rows.Next() {
		var id int
		var firstName string
		var lastName string
		err = rows.Scan(&id, &firstName, &lastName)
		employee := Employee{Id: id, FirstName: firstName, LastName: lastName}
		employees = append(employees, employee)
	}
	return employees
}

func GetEmployee(id int) Employee {
	rows, err := db.Query("SELECT * from EMPLOYEES where Id=$1", id)
	if err != nil {
		log.Fatal(err)
	}

	var employee = Employee{}
	for rows.Next() {
		var id int
		var firstName string
		var lastName string
		err = rows.Scan(&id, &firstName, &lastName)
		employee = Employee{Id: id, FirstName: firstName, LastName: lastName}
	}
	return employee
}

func CreateEmployee(employee Employee) error {
	counter, err := getEmployeeCountForID(employee.Id)
	if err != nil {
		fmt.Println("getEmployeeCountForID failed with errors", err.Error())
		return err
	}

	if counter > 0 {
		return errors.New("Id already exists")
	}

	_, insertErr :=
		db.Exec("INSERT INTO employees (Id, FirstName, LastName) VALUES ($1, $2, $3)",
			employee.Id, employee.FirstName, employee.LastName)

	if insertErr != nil {
		fmt.Println("Create Employee failed with errors", insertErr.Error())
		return errors.New("Create Employee failed")
	}
	fmt.Println("Employee Created")
	return nil
}

func DeleteEmploye(id int) error {
	counter, err := getEmployeeCountForID(id)
	if err != nil {
		fmt.Println("getEmployeeCountForID failed with errors", err.Error())
		return err
	}

	if counter == 0 {
		return errors.New("Id does not exists")
	}

	_, deleteErr := db.Exec("DELETE FROM employees WHERE Id = $1", id)
	if deleteErr != nil {
		fmt.Println("Delete Employee failed with errors", deleteErr.Error())
		return errors.New("Delete Employee failed")
	}

	return nil
}

func UpdateEmployee(employee Employee) error {
	counter, err := getEmployeeCountForID(employee.Id)
	if err != nil {
		fmt.Println("getEmployeeCountForID failed with errors", err.Error())
		return err
	}

	if counter == 0 {
		return errors.New("Id does not exists")
	}

	_, updateErr := db.Exec("UPDATE employees SET FirstName = $1, LastName = $2 WHERE Id = $3", employee.FirstName, employee.LastName, employee.Id)
	if updateErr != nil {
		fmt.Println("Update Employee failed with errors", updateErr.Error())
		return errors.New("Update Employee failed")
	}
	return nil
}

func getEmployeeCountForID(id int) (int, error) {
	var counter int
	err := db.QueryRow(`SELECT count(*) FROM EMPLOYEES where id=$1`, id).Scan(&counter)
	if err != nil {
		fmt.Println("Error in query execution", err.Error())
		return counter, err
	}
	return counter, nil
}
