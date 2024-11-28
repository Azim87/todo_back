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
	port     = 5432
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
		panic(err)
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

	createUserTokenTable := `CREATE TABLE IF NOT EXISTS user_tokens (
    id SERIAL PRIMARY KEY NOT NULL,
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    access_token TEXT NOT NULL,
    refresh_token TEXT NOT NULL,
    access_token_expiry TIMESTAMP NOT NULL,
    refresh_token_expiry TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT unique_user_id UNIQUE (user_id)
)`

	_, err = DB.Exec(createUserTokenTable)

	if err != nil {
		panic("Could not create user token table")
	}
}
