package main

import (
	controller "main/controllers"

	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/products", controller.GetProducts)
	app.Post("/user", controller.CreateUser)
	app.Post("/cart/:mobile_number", controller.CreateCart)
	app.Post("/cart/remove/:mobile_number", controller.RemoveItemCart)
	app.Get("/cart/:mobile_number", controller.GetUserCart)
	app.Post("/order/create/:mobile_number", controller.CreateOrder)
	app.Get("/order/utm", controller.GetOrcerFromUtmSource)
}

func main() {
	// Create new Fiber instance
	app := fiber.New()
	// setup API Routes
	setupRoutes(app)
	app.Listen(":3000")
}
