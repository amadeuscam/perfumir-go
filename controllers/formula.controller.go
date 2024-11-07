package controllers

import (
	"strings"

	"github.com/amadeuscam/perfumir-app/initializers"
	"github.com/amadeuscam/perfumir-app/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreateFormula(c *fiber.Ctx) error {
	var payload *models.FormulaInput
	id := c.Params("fmanagement_id")
	uuid_fmanagement, err := uuid.Parse(id)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid UUID format",
		})
	}
	var fmanagement = models.FormulaManagement{ID: &uuid_fmanagement}

	result := initializers.DB.First(&fmanagement)

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

	formula := models.Formula{
		Name:                payload.Name,
		Status:              payload.Status,
		Version:             payload.Version,
		Source:              payload.Source,
		FormulaManagementID: fmanagement.ID,
	}

	resultManagement := initializers.DB.Create(&formula)

	if resultManagement.Error != nil && strings.Contains(resultManagement.Error.Error(), "duplicate key value violates unique") {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"status": "fail", "message": "Formula Management with that name already exists"})
	} else if resultManagement.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": "Something bad happened"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "data": fiber.Map{"formula": formula}})
}

func GetFormulas(c *fiber.Ctx) error {
	id := c.Params("fmanagement_id")
	uuid, err := uuid.Parse(id)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid UUID format",
		})
	}
	var formulasFmanagement = models.FormulaManagement{ID: &uuid}
	result := initializers.DB.Preload("Formulas").Preload("Formulas.Coments").Preload("Formulas.FormulaIngredients").First(&formulasFmanagement)
	if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": "Something bad happened"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"formulas": formulasFmanagement.Formulas}})
}

func GetFormula(c *fiber.Ctx) error {
	fmanagementId := c.Params("fmanagement_id")
	formulaId := c.Params("formulaid")
	uuid_formula, err := uuid.Parse(formulaId)

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
	var formula models.Formula

	if err := initializers.DB.Where("formula_management_id = ? AND id = ?", uuid_fmanagement, uuid_formula).Preload("Formulas.Coments").First(&formula).Error; err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": "Formula  not found"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"formula": formula}})
}

func UpdateFormula(c *fiber.Ctx) error {
	var payload *models.FormulaInput
	fmanagementId := c.Params("fmanagement_id")
	formulaId := c.Params("formulaid")
	uuid_formula, err := uuid.Parse(formulaId)

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

	var fmanagement = models.FormulaManagement{ID: &uuid_fmanagement}

	result := initializers.DB.First(&fmanagement)

	if result.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": "Something bad happened"})
	}

	var formula models.Formula

	if err := initializers.DB.Where("formula_management_id = ? AND id = ?", fmanagement.ID, uuid_formula).First(&formula).Error; err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": "Formula Management not found"})
	}

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	errors := models.ValidateStruct(payload)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "errors": errors})

	}

	initializers.DB.Model(&formula).Updates(map[string]interface{}{
		"Name":    payload.Name,
		"Status":  payload.Status,
		"Version": payload.Version,
		"Source":  payload.Source,
	})

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"formula": formula}})
}
