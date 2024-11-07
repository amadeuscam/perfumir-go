package routes

import (
	"github.com/amadeuscam/perfumir-app/controllers"
	"github.com/amadeuscam/perfumir-app/middleware"
	"github.com/gofiber/fiber/v2"
)

func ProjectRoutes(app *fiber.App) {
	app.Route("/api/project", func(router fiber.Router) {
		router.Post("/", middleware.DeserializeUser, controllers.CreateProject)
		router.Get("/", middleware.DeserializeUser, controllers.GetAllProjects)
		router.Get("/:id", middleware.DeserializeUser, controllers.GetProject)
		router.Put("/:id", middleware.DeserializeUser, controllers.UpdateProject)
	})

}
