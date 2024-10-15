package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/pred695/auth-microservice/controllers"
	"github.com/pred695/auth-microservice/middleware"
)

func SetUpUserroutes(app *fiber.App) {
	app.Get("/users", controllers.GetUsers)
	app.Post("/login", controllers.LoginUser)
	app.Post("/register", controllers.RegisterUser)

	private := app.Group("/private")
	private.Use(middleware.VerifyUser)
	private.Get("/refresh", controllers.RefreshToken)
	private.Get("/logout", controllers.LogOutUser)
}
