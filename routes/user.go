package routes

import (
	"bookecom/middleware"
	userController "bookecom/controllers/user"
	"github.com/gofiber/fiber/v2"
)

func UserRoutes(group fiber.Router) {

	userGroup := group.Group("/user")

	userGroup.Get("/me", middleware.TokenValidation, userController.GetUserByID)
	userGroup.Post("/me/activate", userController.ActivateUser)
	userGroup.Delete("/me/delete", middleware.TokenValidation, userController.DeleteUser)

}