package routes

import (
	"bookecom/controllers/review"
	"bookecom/middleware"

	"github.com/gofiber/fiber/v2"
)

func ReviewRoutes(group fiber.Router) {

	reviewGroup := group.Group("/review")

	reviewGroup.Post("/", middleware.TokenValidation, review.CreateReviewController)
	reviewGroup.Get("/book/bookId", middleware.TokenValidation, review.GetReviewByBookId)
	reviewGroup.Delete("/:reviewId", middleware.TokenValidation, review.DeleteReview)
}