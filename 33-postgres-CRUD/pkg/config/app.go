package config

import (
	"fmt"
	"log"

	"github.com/jackc/pgx"
)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "postgres"
	DB_NAME     = "GO-SERVER"
	DB_PORT     = "5432"
	DB_HOST     = "localhost"
)

const databaseUrl = "postgres://" + DB_USER + ":" + DB_PASSWORD + "@" + DB_HOST + ":" + DB_PORT + "/" + DB_NAME

var db *pgx.Conn
var connectionError error

func Connect() {
	config, err := pgx.ParseConnectionString(databaseUrl)
	if err != nil {
		fmt.Println("Error parsing connection string:", err.Error())
		log.Fatal(err)
	}
	db, connectionError = pgx.Connect(config)
	if connectionError != nil {
		fmt.Println("Error establishing connection:", connectionError.Error())
		log.Fatal(err)
	}

	fmt.Println("Postgres connection established")

}

func CreateTables() {
	createEmployeeTable()
}

func createEmployeeTable() {
	_, err :=
		db.Exec(
			`CREATE TABLE IF NOT EXISTS EMPLOYEES(
					Id int Primary Key,
					FirstName varchar(30) not null,
					LastName varchar(30) not null
				);`)

	if err != nil {
		fmt.Println("Employee table creation failed:", err.Error())
		log.Fatal(err)
	}

	fmt.Println("Employees table created")
}

func GetDB() *pgx.Conn {
	return db
}
