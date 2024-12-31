package cart

import (
	"bookecom/database"
	"bookecom/models"
	"bookecom/service"
	"fmt"

	cartSchema "bookecom/schemas/cart"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func AddItemToCart(c *fiber.Ctx) error {

	userID := c.Locals("userID").(uuid.UUID)

	user, err := service.GetUserById(userID)
	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "fail",
			"message": "Failed to get userdetails",
		})
	}

	var payload cartSchema.AddItemToCartSchema
	if err := c.BodyParser(&payload); err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "fail",
			"message": "Invalid request body",
		})
	}

	if payload.Quantity < 1 || payload.Quantity > 5{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "fail",
			"message": "Quantity should be between 1 and 5",
		})
	}

	if user.CartId.String() == "00000000-0000-0000-0000-000000000000" || user.CartId == uuid.Nil{
		cart, err := service.CreateCart(user.ID)
		if err != nil{
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status": "fail",
				"message": "Failed to create cart",
			})
		}

		user.CartId = cart.ID

		err = database.DB.Save(&user).Error
		if err != nil{
			return err
		}
	}

	activeCart, err := service.GetCartById(user.CartId)
	if err != nil{
		return err
	}

	if activeCart.Active == false{
		cart, _ := service.CreateCart(userID)
		user.CartId = cart.ID

		database.DB.Save(&user)
	}

	var cart models.Cart
	result := database.DB.Where("id = ?", user.CartId).Preload("Items").First(&cart)
	if result.Error != nil{
		return result.Error
	}

	for _, item := range cart.Items{
		if item.BookID == payload.BookID{
			fmt.Println(item.Quantity)
			item.Quantity = payload.Quantity
			fmt.Println(item.Quantity)
			database.DB.Save(&item)
			return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": "Item quantity updated"})
		}
	}

	cartItem := models.CartItem{
		BookID: payload.BookID,
		Quantity: payload.Quantity,
		CartID: user.CartId,
	}

	cart.Items = append(cart.Items, cartItem)

	err = database.DB.Save(&cart).Error
	if err != nil {
		fmt.Println("Error saving cart:", err)
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "data": cart})

}