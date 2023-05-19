package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var Db *sql.DB

func ConnectToDatabase() {
	dsn := fmt.Sprintf("host=db user=%s password=%s dbname=%s port=5432 sslmode=disable TimeZone=Utc", 
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"), 
		os.Getenv("DB_NAME"),
	)

	db, err := sql.Open("postgres", dsn)

	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	log.Println("Connected to database")

	Db = db
}