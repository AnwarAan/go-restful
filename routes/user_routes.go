package routes

import (
	"fiber-mongo-api/controllers"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App) {
	app.Get("/users", controllers.GetUser)
	app.Get("/user/:userId", controllers.GetUser)
	app.Post("/user", controllers.CreateUser)
	app.Put("/user/:userId", controllers.UpdateUser)
	app.Delete("/user/:userId", controllers.DeleteUser)
}
