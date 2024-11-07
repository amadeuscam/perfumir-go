package controllers

import (
	"github.com/amadeuscam/perfumir-app/initializers"
	"github.com/amadeuscam/perfumir-app/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateFormulaIngredient(c *fiber.Ctx) error {
	var payload *models.FormulaIngredientInput
	id := c.Params("formula_id")
	uuid_formula, err := uuid.Parse(id)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid UUID format",
		})
	}

	var formula = models.Formula{ID: &uuid_formula}

	result := initializers.DB.First(&formula)

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

	formula_ingredient_new := models.FormulaIngredient{
		Name:      payload.Name,
		Amount:    payload.Amount,
		Alcohol:   payload.Alcohol,
		Dilution:  payload.Dilution,
		FormulaID: formula.ID,
	}

	result_create := initializers.DB.Create(&formula_ingredient_new)

	if result_create.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": "Something bad happened"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "data": fiber.Map{"formulaIngredient": formula_ingredient_new}})
}

func DeleteFormulaIngredient(c *fiber.Ctx) error {
	id := c.Params("formula_ingredient_id")

	// Validate UUID format
	uuid, err := uuid.Parse(id)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid UUID format",
		})
	}

	// Fetch the formula ingredient by ID
	var formulaIngredient = models.FormulaIngredient{ID: &uuid}
	result := initializers.DB.First(&formulaIngredient)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"status": "error", "message": "Formula ingredient not found"})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Something bad happened"})
	}

	// Check if the formula is still referenced
	var count int64
	initializers.DB.Model(&models.Formula{}).Where("formula_id = ?", formulaIngredient.FormulaID).Count(&count)

	if count > 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Formula ingredient cannot be deleted as it is still referenced by a formula"})
	}

	// Delete the formula ingredient
	resultDelete := initializers.DB.Delete(&formulaIngredient)

	if resultDelete.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Something bad happened"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Formula ingredient deleted successfully"})
}


// func DeleteAddress(c *fiber.Ctx) error {
//     var address models.FormulaIngredient
//     if err := c.BodyParser(&address); err != nil {
//         return fiber.NewError( fiber.StatusUnprocessableEntity, "Invalid request body")
//     }

//     db := initializers.DB.Delete(&address)
//     if db.Error != nil {
//         return fiber.NewError(fiber.StatusInternalServerError, "Failed to delete address")
//     }

//     return c.SendStatus(fiber.StatusNoContent)
// }
