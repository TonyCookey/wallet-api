package main

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"wallet-api/database"
	"wallet-api/routes"
)

func main() {
	// load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// connect to MYSQL databases
	database.ConnectDB(os.Getenv("DATABASE_URL"))

	database.InitializeRedisInstance(os.Getenv("REDIS_URL"), os.Getenv("REDIS_PASSWORD"), 0)

	// initialize api routes
	routes.InitializeRoutes()

}
