package endpoints

import (
	"github.com/DANCANKARANI/tyson/route/bikes"
	"github.com/DANCANKARANI/tyson/route/providers"
	"github.com/DANCANKARANI/tyson/route/users"
	"github.com/DANCANKARANI/tyson/utilities"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func CreateEndpoint() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // Allow all origins, change this to specific origins in production
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization", 
	}))
	utilities.StaticFileMiddleware(app)
	users.SetUserRoutes(app)
	bikes.SetBikeRoutes(app)
	providers.SetProviderRoutes(app)
	//port
	app.Listen(":3000")
}