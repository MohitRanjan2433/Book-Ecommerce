package review

import (
	reviewSchema "bookecom/schemas/review"
	"bookecom/service"

	"github.com/gofiber/fiber"
	"github.com/google/uuid"
)

func CreateReviewController(c *fiber.Ctx) error {

	var payload reviewSchema.CreateReviewSchema

	userId := c.Locals("userID").(uuid.UUID)
	user, err := service.GetUserById(userId)
	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "fail",
			"message": "Error while fetching userDetails",
		})
	}

	userName := user.Username
	if err := c.BodyParser(&payload); err != nil{
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "fail",
			"message": "Invalid request body",
		})
	}

	results, err := service.CreateReview(userId, payload.BookID, userName, payload)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "fail",
			"message": "Error while creating the review",
		})
	}

	// Return success response with created review details
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  "success",
		"message": "Review created successfully",
		"data":    results,
	})
	

}