package routes

import (
	"github.com/amadeuscam/perfumir-app/controllers"
	"github.com/amadeuscam/perfumir-app/middleware"
	"github.com/gofiber/fiber/v2"
)

func CategoryRoutes(app *fiber.App) {
	app.Route("/api/categorys", func(router fiber.Router) {
		router.Post("/", middleware.DeserializeUser, controllers.CreateCategory)
		router.Delete("/:id", middleware.DeserializeUser, controllers.DeleteCategory)
		router.Get("/", middleware.DeserializeUser, controllers.GetAllCategory)
		router.Get("/:id", middleware.DeserializeUser, controllers.GetCategory)
		router.Put("/:id", middleware.DeserializeUser, controllers.UpdatecCategory)
	})

}
