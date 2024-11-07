package controllers

import (
	"fmt"
	"strings"

	"github.com/amadeuscam/perfumir-app/initializers"
	"github.com/amadeuscam/perfumir-app/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreateFormulaManagement(c *fiber.Ctx) error {
	var payload *models.FormulaManagementInput
	id := c.Params("project_id")
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

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	errors := models.ValidateStruct(payload)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "errors": errors})

	}

	fmanagement := models.FormulaManagement{
		Name:      payload.Name,
		Status:    payload.Status,
		Version:   payload.Version,
		ProjectID: project.ID,
	}

	resultManagement := initializers.DB.Create(&fmanagement)

	if resultManagement.Error != nil && strings.Contains(resultManagement.Error.Error(), "duplicate key value violates unique") {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"status": "fail", "message": "Formula Management with that name already exists"})
	} else if resultManagement.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": "Something bad happened"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "data": fiber.Map{"project": fmanagement}})
}

func UpdateFormulaManagement(c *fiber.Ctx) error {
	var payload *models.FormulaManagementInput
	projectId := c.Params("project_id")
	fmanagementId := c.Params("fmanagement_id")
	uuid_project, err := uuid.Parse(projectId)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid UUID format",
		})
	}
	uuid_fmanagement, err := uuid.Parse(fmanagementId)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid UUID format",
		})
	}

	var project = models.Project{ID: &uuid_project}

	result := initializers.DB.Preload("FormulasManagement").First(&project)

	if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": "Something bad happened"})
	}

	var fmanagement models.FormulaManagement

	if err := initializers.DB.Where("project_id = ? AND id = ?", project.ID, uuid_fmanagement).First(&fmanagement).Error; err != nil {
		fmt.Println("Error: Post no encontrado", err)
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": "Formula Management not found"})
	}

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	errors := models.ValidateStruct(payload)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "errors": errors})

	}

	initializers.DB.Model(&fmanagement).Updates(map[string]interface{}{
		"Name":    payload.Name,
		"Status":  payload.Status,
		"Version": payload.Version,
	})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"fmanagement": fmanagement}})
}
