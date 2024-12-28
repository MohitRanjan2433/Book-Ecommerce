package auth

import (
	userSchema "bookecom/schemas/user"
	"bookecom/service"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func LoginUser(c *fiber.Ctx) error {
	var payload userSchema.LoginUserSchema

	if err := c.BodyParser(&payload); err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "fail",
			"message": "Invalid request payload",
		})
	}

	if payload.Email == ""{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  "fail",
			"message": "Email is required",
		})
	}

	user, err := service.GetUserByEmail(payload.Email)
	if err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Cant find user with this email"})
	}
	

	authResponse, err := service.LoginUser(&payload)
	if err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Invalid credentials"})
	}

	if user.Verified == false{
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "Please verify your account"})
	}

	fmt.Println(authResponse)
	if err != nil{
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "Invalid credentials"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data": fiber.Map{
			"authResponse": authResponse,
		},
	})
}