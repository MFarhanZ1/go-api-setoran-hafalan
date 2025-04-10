package main

import (
    "log"
    "github.com/gofiber/fiber/v3"
)

type Response struct {
	Message string `json:"message"`
}

func main() {
    // Initialize a new Fiber app
    app := fiber.New()

    // Define a route for the GET method on the root path '/'
    app.Get("/", func(c fiber.Ctx) error {
        // Send a json response to the client
        return c.Status(fiber.StatusOK).JSON(Response{
			Message: "Hello, World 👋!",
		})
    })

    // Start the server on port 3000
    log.Fatal(app.Listen(":3000"))
}