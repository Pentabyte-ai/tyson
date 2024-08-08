package providers

import (
	"github.com/DANCANKARANI/tyson/controller/provider"
	"github.com/gofiber/fiber/v2"
)

func SetProviderRoutes(app *fiber.App) {
	auth := app.Group("/api/v1/provider")
	auth.Post("/",provider.CreateProviderAccount)
	auth.Post("/login",provider.Login)
	//protected routes
}