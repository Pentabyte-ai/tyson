package users

import (
	"github.com/DANCANKARANI/tyson/controller/user"
	"github.com/gofiber/fiber/v2"
)

func SetUserRoutes(app *fiber.App) {
	auth := app.Group("/api/v1/user")
	auth.Post("/signup",user.CreateUserAccount)
	auth.Post("/login",user.Login)
	//protected routes
	userGroup := auth.Group("/",user.JWTMiddleware)
	userGroup.Get("/",user.GetUserByIdHandler)
}