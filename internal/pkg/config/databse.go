package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	var err error

	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dbConfig := fmt.Sprintf("user=%s dbname=%s sslmode=%s password=%s", os.Getenv("DB_USER"), os.Getenv("DB_NAME"), os.Getenv("DB_SSL_MODE"), os.Getenv("DB_PASS"))

	DB, err = sql.Open("postgres", dbConfig)
	if err != nil {
		log.Fatalf("Error opening database: %q", err)
	}
	if err = DB.Ping(); err != nil {
		log.Fatalf("Error connecting to the database: %q", err)
	}
}
