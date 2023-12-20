package routes

import (
	c "go-day2/controllers"

	"github.com/gofiber/fiber/v2"
	_ "github.com/gofiber/fiber/v2/middleware/basicauth"
)

func InetRoutes(app *fiber.App) {
	api := app.Group("/api")
	v1 := api.Group("/v1")
	v3 := api.Group("/v3")

	v1.Post("/register", c.Register)
	v3.Get("/fact/:number", c.FactParams)
	v3.Get("/tan", c.AsciiConv)

}
