package user

import (
	"bookecom/service"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)


func GetAllUsers(c *fiber.Ctx) error {

	user, err := service.GetAllUsers()
	if err != nil {
		// Handle error
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": "Failed to get user details"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": user})
}

func GetUserByID(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uuid.UUID)

	user, err := service.GetUserById(userID)
	if err != nil {
		// Handle error
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": "Failed to get user details"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": user})
}