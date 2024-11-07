package controllers

import (
	"github.com/amadeuscam/perfumir-app/initializers"
	"github.com/amadeuscam/perfumir-app/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreateComment(c *fiber.Ctx) error {
	var payload *models.ComentInput

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

	newComment := models.Coment{
		Body:      payload.Body,
		FormulaID: formula.ID,
	}

	result_create := initializers.DB.Create(&newComment)

	if result_create.Error != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": "Something bad happened"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"status": "success", "data": fiber.Map{"comment": newComment}})
}

func DeleteComment(c *fiber.Ctx) error {

	// Get the formula ID and comment ID from URL parameters.
	id := c.Params("formula_id")
	uuid_formula, err := uuid.Parse(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid UUID format for formula ID",
		})
	}

	commentId := c.Params("comment_id")
	uuid_comment, err := uuid.Parse(commentId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid UUID format for comment ID",
		})
	}

	// Define a Comment struct to interact with the database.
	var comment models.Coment

	// Delete the comment from the database based on formula_id and id.
	if err := initializers.DB.Where("formula_id = ? AND id = ?", uuid_formula, uuid_comment).Delete(&comment).Error; err != nil {
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"status": "error", "message": "Something bad happened while deleting the comment"})
	}

	// Return a success response after deletion.
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Comment deleted successfully"})

}
