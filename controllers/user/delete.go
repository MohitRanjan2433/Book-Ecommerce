package user

import (
	"bookecom/service"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)


func DeleteUser(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uuid.UUID)

	user, err := service.GetUserById(userID)
	if err != nil {
		// Handle error
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": "Failed to get user details"})
	}

	user.Active = false
	_, err = service.UpdateUser(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": "Failed to delete user"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "User deleted successfully"})
}