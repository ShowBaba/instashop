package router

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"instashop/internal/routes"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to InstaShop API")
}

func Routes(app *fiber.App, database *gorm.DB) {
	apiURL := "/"
	router := app.Group(apiURL)

	app.Get(apiURL, welcome)
	routes.RegisterAuthRoutes(router, database)
	routes.RegisterUserRoutes(router, database)
	routes.RegisterOrderRoutes(router, database)
	routes.RegisterProductRoutes(router, database)
}
