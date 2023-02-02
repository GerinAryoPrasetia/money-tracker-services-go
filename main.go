package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env files")
	}

	dbName := os.Getenv("DB_NAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbUsername := os.Getenv("DB_USERNAME")

	conn := fmt.Sprintf("postgresql://%s:%s@localhost:5432/%s?sslmode=disable", dbUsername, dbPassword, dbName)
	db, err := sql.Open("postgres", conn)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
}
