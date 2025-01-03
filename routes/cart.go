package routes

import (
	cartController "bookecom/controllers/cart"
	"bookecom/middleware"

	"github.com/gofiber/fiber/v2"
)

func CartRoutes(group fiber.Router){

	cartGroup := group.Group("/cart")

	cartGroup.Post("/items", middleware.TokenValidation, cartController.AddItemToCart)
	cartGroup.Delete("/items", middleware.TokenValidation, cartController.DeleteCart)
	cartGroup.Get("/", middleware.TokenValidation, cartController.GetUserCart)
	cartGroup.Get("/all", cartController.GetAllCarts)
}