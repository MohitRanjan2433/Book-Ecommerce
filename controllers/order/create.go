package order

import (
	"bookecom/database"
	"bookecom/models"
	"bookecom/service"
	"time"

	orderSchema "bookecom/schemas/order"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func CreateOrderController(c *fiber.Ctx) error {

	userId := c.Locals("userID").(uuid.UUID)
	user, err := service.GetUserById(userId)
	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "fail",
			"message": "failed to fetch userdetails",
		})
	}

	var payload orderSchema.CreateOrderSchema
	if err := c.BodyParser(&payload); err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "fail",
			"message": "invalid request body",
		})
	}

	if user.CartId.String() == "00000000-0000-0000-0000-000000000000" || user.CartId == uuid.Nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "fail",
			"message": "User does not have a cart",
		})
	}

	var cart models.Cart
	result := database.DB.Preload("Items").Where("user_id = ?", userId).First(&cart)
	if result.Error != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "fail",
			"message": "failed to fetch cart",
		})
	}

	//cal total price

	cost := 0.0	
	for _, item := range cart.Items{
		var book models.Book
		result := database.DB.First(&book, "id = ?",item.BookID)
		if result.Error != nil{
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status": "fail",
				"message": "failed to fetch book details",
			})
		}

		cost = cost + (book.Price * float64(book.Quantity))
	}

	result = database.DB.Preload("Items").Find(&cart, "id = ?", user.CartId)
	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": "Failed to get carts"})
	}

	// check if those many items are even available or not in inventory else tell the user to reduce the quantity
	for _, item := range cart.Items {
		book, err := service.GetBookById(item.BookID)
		if err != nil{
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"status": "fail",
				"message": "failed to fetch book details",
			})
		}

		if(book.Quantity < item.Quantity){
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"status": "fail",
				"message": "Not enough quantity of the book in stock",
			})
		}
	}

	now := time.Now()
	order := models.Order{
		ID:              uuid.New(),
		CartID:          payload.CartID,
		TotalCost:       cost,
		CreatedAt:       now,
		UpdatedAt:       now,
	}

	database.DB.Save(&order)

	user.Orders = append(user.Orders, order.ID.String())

	// Save the updated user record
	err = database.DB.Save(&user).Error
	if err != nil {
		return err
	}

	cart.Active = false

	err = database.DB.Save(&cart).Error
	if err != nil {
		return err
	}

	user, err = service.GetUserById(userId)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "fail", "message": "Failed to get user details"})
	}

	user.CartId = uuid.Nil
	database.DB.Save(&user)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"status": "success", "message": "Order created successfully", "order": order})
}