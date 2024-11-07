package controllers

import (
	"github.com/amadeuscam/perfumir-app/initializers"
	"github.com/amadeuscam/perfumir-app/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// Get all ingredients
func GetAllIngredients(c *fiber.Ctx) error {
	var ingredients []models.Ingredient
	if err := initializers.DB.Find(&ingredients).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"ingredients": ingredients}})
}

// Get a single ingredient by ID
func GetIngredient(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid UUID"})
	}
	var ingredient models.Ingredient
	if err := initializers.DB.First(&ingredient, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Ingredient not found"})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"ingredient": ingredient}})
}

// Create a new ingredient
func CreateIngredient(c *fiber.Ctx) error {
	var payload *models.Ingredient

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}

	if err := initializers.DB.Create(&payload).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"ingredient": payload}})
}

// Update an ingredient
func UpdateIngredient(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid UUID"})
	}
	var ingredient models.Ingredient
	if err := initializers.DB.First(&ingredient, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Ingredient not found"})
	}
	if err := c.BodyParser(&ingredient); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}
	if err := initializers.DB.Save(&ingredient).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": fiber.Map{"ingredient": ingredient}})
}

// Delete an ingredient
func DeleteIngredient(c *fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid UUID"})
	}
	var ingredient models.Ingredient
	if err := initializers.DB.First(&ingredient, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Ingredient not found"})
	}
	if err := initializers.DB.Delete(&ingredient).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "Ingredient deleted successfully"})

}
