package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")

	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, password, host, port, name)
	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		log.Fatal("Failed to connect to the database: ", err)
	} else {
		DB = db
		fmt.Println("Database successfully connected.")
		fmt.Printf("%+v", db)
	}
}

func CloseDB() {
	if err := DB.Close(); err != nil {
		log.Println("Error closing the database: ", err)
	}
}
