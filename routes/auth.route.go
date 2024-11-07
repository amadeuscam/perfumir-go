package routes

import (
	"github.com/amadeuscam/perfumir-app/controllers"
	"github.com/amadeuscam/perfumir-app/middleware"
	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app *fiber.App) {
	app.Route("/auth", func(router fiber.Router) {
		router.Post("/register", controllers.SignUpUser)
		router.Post("/login", controllers.SignInUser)
		router.Get("/logout", middleware.DeserializeUser, controllers.LogoutUser)
		router.Get("/users/me", middleware.DeserializeUser, controllers.GetMe)
	})

}
