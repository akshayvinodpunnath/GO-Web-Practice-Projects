/*
databaseUrl := "postgres://postgres:mypassword@localhost:5432/postgres"

// this returns connection pool
dbPool, err := pgxpool.Connect(context.Background(), databaseUrl)

*/

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jackc/pgx"
)

const (
	CONNECT_HOST = "localhost"
	CONNECT_PORT = "8080"
	DB_USER      = "postgres"
	DB_PASSWORD  = "postgres"
	DB_NAME      = "GO-SERVER"
	DB_PORT      = "5432"
	DB_HOST      = "localhost"
)

const databaseUrl = "postgres://" + DB_USER + ":" + DB_PASSWORD + "@" + DB_HOST + ":" + DB_PORT + "/" + DB_NAME

var connectionError error
var db *pgx.Conn

func init() {
	config, err := pgx.ParseConnectionString(databaseUrl)
	if err != nil {
		log.Fatal(err)
	}
	db, connectionError = pgx.Connect(config)
	if connectionError != nil {
		log.Fatal(connectionError)
	}
	fmt.Println("Postgres connection aquired")
}

func getCurrentDb(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT current_database() as dbname")
	if err != nil {
		log.Fatal(err)
	}
	var dbname string
	for rows.Next() {
		rows.Scan(&dbname)
	}
	fmt.Fprintf(w, "Current Database is :: %s", dbname)
}

func main() {
	http.HandleFunc("/", getCurrentDb)

	fmt.Println("Starting server on port:", CONNECT_PORT)
	log.Fatal(http.ListenAndServe(CONNECT_HOST+":"+CONNECT_PORT, nil))
}
