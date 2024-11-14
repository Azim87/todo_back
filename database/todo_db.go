package database

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

const (
	host     = "localhost"
	port     = 5544
	user     = "postgres"
	password = "123456"
	dbname   = "tododb"
)

func Connect() {
	var err error

	connStr := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	DB, err = sqlx.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTable()

	log.Println("Connected to database...")
}

func createTable() {
	createTodoTable := `
	CREATE TABLE IF NOT EXISTS todo (
    	id SERIAL PRIMARY KEY NOT NULL,
    	title VARCHAR(255),
    	description TEXT,
    	completed BOOLEAN
	)`
	_, err := DB.Exec(createTodoTable)

	if err != nil {
		panic("Could not create todo table")
	}

	createUserTable := `
	CREATE TABLE IF NOT EXISTS users (
    	id SERIAL PRIMARY KEY NOT NULL,
    	email VARCHAR(255),
    	secured_password VARCHAR(255),
		password VARCHAR(255)
	)`
	_, err = DB.Exec(createUserTable)

	if err != nil {
		panic("Could not create user table")
	}
}
