package main

import (
	"fmt"
	"log"
	"os"

	"github.com/amadeuscam/perfumir-app/initializers"
	"github.com/amadeuscam/perfumir-app/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatalln("Failed to load environment variables! \n", err.Error())
	}
	initializers.ConnectDB(&config)
}

func main() {
	app := fiber.New()

	// micro := fiber.New()

	// app.Mount("/api", micro)

	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowMethods:     "GET, POST",
		AllowCredentials: true,
	}))

	app.Use(logger.New(logger.Config{
		Format:     "[${time}] ${status} - ${method} ${path}\n", // Formato personalizado
		TimeFormat: "02-Jan-2006",
		TimeZone:   "Local",
		Output:     os.Stdout, // Puedes cambiar la salida, como escribir en un archivo
	}))

	app.Get("/healthchecker", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"status":  "success",
			"message": "Welcome to Golang, Fiber, and GORM",
		})
	})

	routes.AuthRoutes(app)
	routes.ProjectRoutes(app)
	routes.FmanagementRoutes(app)
	routes.FormulaRoutes(app)
	routes.CommentRoutes(app)
	routes.FormulaIngredientsRoutes(app)
	routes.IngredientRoutes(app)
	routes.CategoryRoutes(app)

	app.All("*", func(c *fiber.Ctx) error {
		path := c.Path()
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "fail",
			"message": fmt.Sprintf("Path: %v does not exists on this server", path),
		})
	})

	log.Fatal(app.Listen(":8000"))
}
