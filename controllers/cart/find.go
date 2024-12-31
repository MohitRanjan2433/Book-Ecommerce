package cart

import (
	"bookecom/database"
	"bookecom/models"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetAllCarts(c *fiber.Ctx) error {
	var carts []models.Cart

	result := database.DB.Preload("Items").Find(&carts)
	if result.Error != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "fail",
			"message": "failed to fetch the cart items",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"message": "Cart fetched successfully",
		"data": carts,
	})
}

func GetUserCart(c *fiber.Ctx) error {
	userID := c.Locals("userID").(uuid.UUID)

	var cart models.Cart

	result:= database.DB.Preload("Items").Where("user_id = ?", userID).First(&cart)
	if result.Error != nil{
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status": "fail",
			"message": "Cart not found",
		})
	}

	cost := 0.0

	for _, item := range cart.Items{
		var book models.Book
		result := database.DB.First(&book, "id = ?", item.BookID)
		if result.Error != nil{
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status": "fail",
				"message": "Fail to fetch BookDetails from the cart items",
			})
		}

		cost = cost + (book.Price * float64(book.Quantity))
	}

	cart.TotalCost = cost

	err := database.DB.Save(&cart).Error
	if err != nil {
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": cart})
}