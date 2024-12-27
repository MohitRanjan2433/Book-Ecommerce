package middleware

import (
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"
	"github.com/google/uuid"
)

func TokenValidation(c *fiber.Ctx) error {
	secretKey := os.Getenv("ACCESS_TOKEN_SECRET")

	authHeader := c.Get("Authorization")
	if authHeader == ""{
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Unauthorized",
			"message": "Authorization header missing"
		})
	}

	accessToken := authHeader[len("Bearer "):]
	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(accessToken, claims, func (token *jwt.Token) (interface{}, error)  {
		return []byte(secretKey), nil
	})

	if err != nil || !token.Valid{
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":"fail",
			"message":"Invalid accesstoken or expired",
		})
	}

	userIDStr, ok := claims["userID"].(string)
	role, ok := claims["role"].(string)
	if !ok{
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "Invalid access token"})
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil{
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": "Invalid access token"})
	}

	c.Locals("userID", userID)
	c.Locals("role", role)

	return c.Next()
}