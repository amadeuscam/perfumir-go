package routes

import (
	"github.com/amadeuscam/perfumir-app/controllers"
	"github.com/amadeuscam/perfumir-app/middleware"
	"github.com/gofiber/fiber/v2"
)

func CommentRoutes(app *fiber.App) {
	app.Route("/api/comments", func(router fiber.Router) {
		router.Post("/:formula_id", middleware.DeserializeUser, controllers.CreateComment)
		router.Delete("/:formula_id/:comment_id", middleware.DeserializeUser, controllers.DeleteComment)
		// router.Get("/:fmanagement_id", middleware.DeserializeUser, controllers.GetFormulas)
		// router.Put("/:fmanagement_id/:formulaid", middleware.DeserializeUser, controllers.UpdateFormula)
	})

}
