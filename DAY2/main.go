package main

import (
	"go-day2/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// character to ASCII
	app := fiber.New()
	routes.InetRoutes(app)
	app.Listen(":3000")

}
