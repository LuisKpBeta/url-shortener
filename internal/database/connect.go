package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	HOST          = "localhost"
	DATABASE      = "url"
	USER          = "postgres"
	PASSWORD      = "postgres"
	MAX_OPEN_CONN = 10
	MAX_IDLE_CONN = 10
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func ConnectToDatabase() *sql.DB {
	var connectionString string = fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", HOST, USER, PASSWORD, DATABASE)
	db, err := sql.Open("postgres", connectionString)
	checkError(err)

	err = db.Ping()
	checkError(err)
	db.SetMaxOpenConns(MAX_OPEN_CONN)
	db.SetMaxOpenConns(MAX_IDLE_CONN)

	return db
}
