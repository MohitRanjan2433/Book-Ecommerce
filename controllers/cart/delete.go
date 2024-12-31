package cart

import (
	"bookecom/database"
	"bookecom/service"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func DeleteCart(c *fiber.Ctx) error{

	userID := c.Locals("userID").(uuid.UUID)
	
	err := service.DeleteUserCart(userID)
	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "fail",
			"message": "Failed to delete cart",
		})
	}

	user, err := service.GetUserById(userID)
	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "fail",
			"message": "Failed to get user",
		})
	}

	user.CartId = uuid.Nil
	database.DB.Save(&user)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"message": "Cart deleted successfully",
	})
}