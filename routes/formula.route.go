package routes

import (
	"github.com/amadeuscam/perfumir-app/controllers"
	"github.com/amadeuscam/perfumir-app/middleware"
	"github.com/gofiber/fiber/v2"
)

func FormulaRoutes(app *fiber.App) {
	app.Route("/api/formula", func(router fiber.Router) {
		router.Post("/:fmanagement_id", middleware.DeserializeUser, controllers.CreateFormula)
		router.Get("/:fmanagement_id/:formulaid", middleware.DeserializeUser, controllers.GetFormula)
		router.Get("/:fmanagement_id", middleware.DeserializeUser, controllers.GetFormulas)
		router.Put("/:fmanagement_id/:formulaid", middleware.DeserializeUser, controllers.UpdateFormula)
	})

}
