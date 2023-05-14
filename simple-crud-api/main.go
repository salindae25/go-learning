package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/salindae/simple-crud-api/controllers"
	"github.com/salindae/simple-crud-api/initializers"
)

func init() {
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}
func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Route("/product", controllers.ProductRoutes)

	app.Listen(":3000")
}
