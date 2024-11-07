package routes

import (
	"github.com/amadeuscam/perfumir-app/controllers"
	"github.com/amadeuscam/perfumir-app/middleware"
	"github.com/gofiber/fiber/v2"
)

func FormulaIngredientsRoutes(app *fiber.App) {
	app.Route("/api/formula-ingredient", func(router fiber.Router) {
		router.Post("/:formula_id", middleware.DeserializeUser, controllers.CreateFormulaIngredient)
		router.Delete("/:formula_ingredient_id", middleware.DeserializeUser, controllers.DeleteFormulaIngredient)
		// router.Get("/:fmanagement_id", middleware.DeserializeUser, controllers.GetFormulas)
		// router.Put("/:fmanagement_id/:formulaid", middleware.DeserializeUser, controllers.UpdateFormula)
	})

}
