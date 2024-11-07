package controllers

import (
	"strings"

	"github.com/amadeuscam/perfumir-app/initializers"
	"github.com/amadeuscam/perfumir-app/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreateProject(c *fiber.Ctx) error {
	var payload *models.ProjectInput

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	errors := models.ValidateStruct(payload)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "errors": errors})

	}

	newProject := models.Project{
		Name: payload.Name,
	}

	result := initializers.DB.Create(&newProject)

	if result.Error != nil && strings.Contains(result.Error.Error(), "duplicate key value violates unique") {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"status": "fail", "message": "Project with that name already exists"})
	} else if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": "Something bad happened"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "data": fiber.Map{"project": newProject}})
}

func GetAllProjects(c *fiber.Ctx) error {
	var projects = []models.Project{}
	// Get all records
	result := initializers.DB.Preload("FormulasManagement").Find(&projects)
	if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": "Something bad happened"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"projects": projects}})
}

func GetProject(c *fiber.Ctx) error {
	id := c.Params("id")
	uuid, err := uuid.Parse(id)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid UUID format",
		})
	}
	var project = models.Project{ID: &uuid}
	result := initializers.DB.Preload("FormulasManagement").First(&project)
	if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": "Something bad happened"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"project": project}})
}

func UpdateProject(c *fiber.Ctx) error {
	id := c.Params("id")
	uuid, err := uuid.Parse(id)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid UUID format",
		})
	}

	var project = models.Project{ID: &uuid}

	result := initializers.DB.First(&project)
	if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": "Something bad happened"})
	}

	var updatedData models.Project

	// Parsear los datos JSON del cuerpo de la solicitud en `updatedData`
	if err := c.BodyParser(&updatedData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	// Actualizar los campos del usuario con los datos proporcionados
	if updatedData.Name != "" {
		project.Name = updatedData.Name
	}

	initializers.DB.Model(&project).Update("name", updatedData.Name)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"project": project}})
}
