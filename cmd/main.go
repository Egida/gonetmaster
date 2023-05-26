package main

import (
	"log"

	"github.com/edwardelton/gonetmaster/api/database"
	"github.com/edwardelton/gonetmaster/api/router"
	aggregatedLogger "github.com/edwardelton/gonetmaster/logger"
	"github.com/edwardelton/gonetmaster/util"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	file := util.RetrieveOrCreateFile("./api/logs/logs.txt")

	aggregatedLogger.Log = aggregatedLogger.New(file)

	err := godotenv.Load()

	if err != nil {
		aggregatedLogger.Log.Error("Error loading .env file")
	}

	database.InitializeDatabase()
	defer database.Db.Close()

	app := fiber.New()

	app.Use(logger.New(logger.Config{
		Format:     "${time} | ${method} | ${path} | ${ip} | ${status} | ${latency}\n",
		TimeFormat: "02-Jan-2006 15:04:05",
		Output:     file,
	}))

	router.SetupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}