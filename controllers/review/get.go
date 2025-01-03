package review

import (
	"bookecom/service"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetReviewByBookId(c *fiber.Ctx) error {
	bookId := uuid.MustParse(c.Params("bookId"))	
	
	result, err := service.GetReviewByBookId(bookId)
	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "fail",
			"message": "Error while getting the service call",
		})
	}

	if len(result) == 0{
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status": "fail",
			"message": "No reviews found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"message": "Reviews found",
		"total_reviews": len(result),
		"data": result,
	})

}