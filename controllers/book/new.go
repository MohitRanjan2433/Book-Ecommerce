package book

import (
	"bookecom/service"
	"bookecom/schemas/book"
	"github.com/gofiber/fiber/v2"
)

// CreateBook handles the HTTP request for creating a new book
func CreateBookController(c *fiber.Ctx) error {
	var payload book.CreateBookSchema

	// Parse the incoming JSON request body into the CreateBookSchema struct
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request body: " + err.Error(),
		})
	}

	// Call the service function to create the book
	book, err := service.CreateBook(&payload)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create book: " + err.Error(),
		})
	}

	// Return the created book as a JSON response
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"book": book,
	})
}
