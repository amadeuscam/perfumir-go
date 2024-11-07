package routes

import (
	"github.com/amadeuscam/perfumir-app/controllers"
	"github.com/amadeuscam/perfumir-app/middleware"
	"github.com/gofiber/fiber/v2"
)

func IngredientRoutes(app *fiber.App) {
	app.Route("/api/ingredients", func(router fiber.Router) {
		router.Post("/", middleware.DeserializeUser, controllers.CreateIngredient)
		router.Delete("/:id", middleware.DeserializeUser, controllers.DeleteIngredient)
		router.Get("/", middleware.DeserializeUser, controllers.GetAllIngredients)
		router.Get("/:id", middleware.DeserializeUser, controllers.GetIngredient)
		router.Put("/:id", middleware.DeserializeUser, controllers.UpdateIngredient)
	})

}
