package models

import (
	"32-postgres-create-record/pkg/config"
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

func CreateEmployee(employee Employee) bool {
	var counter int

	err := db.QueryRow(`SELECT count(*) FROM EMPLOYEES where id=$1`, employee.Id).Scan(&counter)
	if err != nil {
		fmt.Println("Error in query execution")
		log.Fatal(err)
	}

	if counter > 0 {
		fmt.Println("we have", counter, "rows")
		return false
	}

	_, insertErr :=
		db.Exec("INSERT INTO employees (Id, FirstName, LastName) VALUES ($1, $2, $3)",
			employee.Id, employee.FirstName, employee.LastName)

	if insertErr != nil {
		fmt.Println("Error in insert")
		log.Fatal(insertErr)
	}

	return true
}
