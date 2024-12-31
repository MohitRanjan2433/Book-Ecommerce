package book

import (
	"bookecom/schemas/book"
	"bookecom/service"

	"github.com/gofiber/fiber/v2"
)


func CreateBookController(c *fiber.Ctx) error {
	var payload book.CreateBookSchema

	userRole := c.Locals("role").(string)

	if userRole == "user"{
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status": "fail",
			"message": "User is not a vendor, only vendors can create Book",
		})
	}

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body: " + err.Error(),
		})
	}


	book, err := service.CreateBook(&payload)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create book: " + err.Error(),
		})
	}


	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"book": book,
	})
}
