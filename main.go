package main

import (
	controller "main/controllers"

	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/products", controller.GetProducts)
}

func main() {
	// Create new Fiber instance
	app := fiber.New()
	// setup API Routes
	setupRoutes(app)
	app.Listen(":3000")
}
