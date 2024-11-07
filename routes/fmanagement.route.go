package routes

import (
	"github.com/amadeuscam/perfumir-app/controllers"
	"github.com/amadeuscam/perfumir-app/middleware"
	"github.com/gofiber/fiber/v2"
)

func FmanagementRoutes(app *fiber.App) {
	app.Route("/api/fmanagement", func(router fiber.Router) {
		router.Post("/:project_id", middleware.DeserializeUser, controllers.CreateFormulaManagement)
		router.Put("/:project_id/:fmanagement_id", middleware.DeserializeUser, controllers.UpdateFormulaManagement)
	})

}
