package main

import (
	"log"
	"os"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/mfarhanz1/go-api-setoran-hafalan/app/routes"
)

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Initialize a new Fiber app
	app := fiber.New()

	// Construct all routes
	app.Use(cors.New())
	app.Use(logger.New())
	routes.GlobalRoute(app)

	// Start the server on port specified in the environment variable APP_PORT
	log.Fatal(app.Listen(":" + os.Getenv("APP_PORT")))
}
