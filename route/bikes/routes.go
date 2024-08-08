package bikes

import (
	"github.com/DANCANKARANI/tyson/controller/bike"
	"github.com/DANCANKARANI/tyson/controller/booking"
	"github.com/DANCANKARANI/tyson/controller/user"
	"github.com/gofiber/fiber/v2"
)

func SetBikeRoutes(app *fiber.App) {
	auth := app.Group("/api/v1/bikes")
	//protected routes
	bikeGroup := auth.Group("/",user.JWTMiddleware)
	bikeGroup.Post("/",bike.AddBikeHandler)
	bikeGroup.Patch("/:id",bike.UpdateBikeHandler)
	bikeGroup.Get("/",bike.GetBikeByLocationHandler)
	bikeGroup.Get("/all",bike.GetAllBikesHandler)
	bikeGroup.Post("/bookings/:id",booking.BookingHandler)

}