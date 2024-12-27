package book

import (
	"bookecom/models"
	"bookecom/service"
	bookSchema "bookecom/schemas/book"
	"github.com/gofiber/fiber/v2"
)


func FindAllBooks(c *fiber.Ctx) error {

	title := c.Query("title", "")
	author := c.Query("author", "")

	var books []models.Book

	// Call the service to get books by title or author
	books, err := service.GetBookByTitleOrAuthon(title, author)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	bookResponses := bookSchema.MapBookToResponse(books)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"results": len(books),
		"books": bookResponses,
	})
}