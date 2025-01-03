package order

import (
	"bookecom/service"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

func containsOrderID(orders pq.StringArray, orderID string) bool {

	for _, id := range orders{
		if id == orderID{
			return true
		}
	}
	return false
}

func GetOrders(c *fiber.Ctx) error {
	userId := c.Locals("userID").(uuid.UUID)
	user, err := service.GetUserById(userId)
	if err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "fail",
			"message": "Error in fetching user details",
		})
	}

	orderIDs := user.Orders
	orders, err := service.GetOrders(orderIDs)
	if err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "fail",
			"message": "Error in fetching orders",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"data": orders,
	})
}

func GetOrderByID(c *fiber.Ctx) error {
	userId := c.Locals("userID").(uuid.UUID)
	user, err := service.GetUserById(userId)
	if err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "fail",
			"message": "Error in fetching user details",
		})
	}

	orderID := uuid.MustParse(c.Params("orderId"))

	if !containsOrderID(user.Orders, orderID.String()){
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status": "fail",
			"message": "Order not found",
		})
	}

	order, err := service.GetOrderById(orderID.String())
	if err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "fail",
			"message": "Error in fetching order",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": order})

}