package book

import (
	"bookecom/service"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func DeleteBook(c *fiber.Ctx) error{

	userID := c.Locals("userID").(uuid.UUID)
	bookID := c.Params("bookID")

	err := service.DeleteBook(userID, bookID)
	if err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"message": "Book deleted successfully",
	})
}