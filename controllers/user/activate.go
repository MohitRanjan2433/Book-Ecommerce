package user

import (
	"bookecom/service"
	userSchema "bookecom/schemas/user"
	"github.com/gofiber/fiber/v2"
)


func ActivateUser(c *fiber.Ctx) error {
	var payload userSchema.LoginUserSchema

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": err.Error()})
	}
	user, err := service.GetUserByEmail(payload.Email)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Cant find user with this email"})
	}

	_, err = service.LoginUser(&payload)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Invalid credentials"})
	}

	user.Active = true
	_, err = service.UpdateUser(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": "Failed to delete user"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "User reactivated successfully"})
}