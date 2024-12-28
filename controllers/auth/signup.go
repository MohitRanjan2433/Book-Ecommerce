package auth

import (
	"bookecom/schemas/user"
	"bookecom/service"

	"github.com/gofiber/fiber/v2"
)

func SignUpController(c *fiber.Ctx) error {

	var payload user.RegisterUserSchema

	if err := c.BodyParser(&payload); err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request Payload",
		})
	}

	if payload.Username == "" || payload.Email == "" || payload.Password == "" || payload.PhoneNumber == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Invalid request body"})
	}

	newUser, err := service.SignupUser(&payload)
	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Internal Server Error",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "User created Successfully",
		"data": newUser,
	})
}