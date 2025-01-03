package routes

import (
	order "bookecom/controllers/order"
	"bookecom/middleware"

	"github.com/gofiber/fiber/v2"
)

func OrderRoutes(group fiber.Router) {

	orderGroup := group.Group("/order")

	orderGroup.Get("/", middleware.TokenValidation, order.GetOrders)
	orderGroup.Post("/", middleware.TokenValidation, order.CreateOrderController)

	orderGroup.Route("/:orderId", func(router fiber.Router) {
		router.Get("/", middleware.TokenValidation, order.GetOrderByID)
		// router.Delete("/", middleware.TokenValidation, order.D)
	})
}