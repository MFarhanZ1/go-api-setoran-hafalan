package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/mfarhanz1/go-api-setoran-hafalan/app/handlers"
	"github.com/mfarhanz1/go-api-setoran-hafalan/app/services"
)

func GlobalRoute(app *fiber.App) {
	// Initialize the global service and handler
	service := services.NewGlobalService()
	handler := handlers.NewGlobalHandler(service)
	
	// Define the routes
	app.Get("/", handler.Introduce)
	app.Get("/*", handler.NotFound)
}