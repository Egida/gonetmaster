package database

import (
	"context"
	"database/sql"
	"time"

	"github.com/edwardelton/gonetmaster/api/config"
	"github.com/edwardelton/gonetmaster/logger"
	_ "github.com/lib/pq"
)

var Db *sql.DB

func InitializeDatabase() {
	connectToDatabase()
	createUserTable()
}

func connectToDatabase() {
	db, err := sql.Open("postgres", dsn())

	if err != nil {
		logger.Log.Error("Connecting to database: ", err)
	}

	logger.Log.Info("Connected to database")

	Db = db
}

func createUserTable() {
	ctx, cancelFunc := context.WithTimeout(context.Background(), config.ConnectionTimeout*time.Second)
	defer cancelFunc()

	createUserTableQuery := `
	CREATE TABLE IF NOT EXISTS client (
		id UUID PRIMARY KEY,
		key VARCHAR(255) NOT NULL,
		inserted_at TIMESTAMP NOT NULL DEFAULT NOW()
	)`

	_, err := Db.ExecContext(ctx, createUserTableQuery)

	if err != nil {
		logger.Log.Error("Creating user table: ", err)
	}
}
