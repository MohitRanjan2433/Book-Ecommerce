package routes

import (
	"github.com/gofiber/fiber/v2"
	"bookecom/controllers/auth"
)

func AuthRoutes(group fiber.Router) {

	authGroup := group.Group("/user")

	authGroup.Post("/signup", auth.SignUpController)
	authGroup.Post("/login", auth.LoginUser)
}