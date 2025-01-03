package review

import (
	"bookecom/service"

	"github.com/gofiber/fiber"
	"github.com/google/uuid"
)

func DeleteReview(c *fiber.Ctx) error {

	userId := c.Locals("userID").(uuid.UUID)

	reviewID := uuid.MustParse(c.Params("reviewId"))
	review, err := service.GetReviewById(reviewID)
	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "fail",
			"message": "Error while fetching review details",
		})
	}

	if userId != review.UserID{
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"status": "fail",
			"message": "You are not authorized to delete this review",
		})
	}

	err = service.DeleteReview(reviewID)

	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "fail",
			"message": "Error while deleting review",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"message": "Review deleted successfully",
	})
}