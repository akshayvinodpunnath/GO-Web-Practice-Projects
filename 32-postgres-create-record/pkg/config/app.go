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

var connectionError error
var db *pgx.Conn

func Connect() {
	config, err := pgx.ParseConnectionString(databaseUrl)
	if err != nil {
		log.Fatal(err)
	}
	db, connectionError = pgx.Connect(config)
	if connectionError != nil {
		log.Fatal(connectionError)
	}
	fmt.Println("Postgres connection aquired")

	_, tableCreateError := db.Exec(`
	CREATE TABLE IF NOT EXISTS EMPLOYEES(
		Id int Primary Key,
		FirstName varchar(30) not null,
		LastName varchar(30) not null
	);`)

	if tableCreateError != nil {
		log.Fatal(tableCreateError)
	}
}

func GetDB() *pgx.Conn {
	return db
}
