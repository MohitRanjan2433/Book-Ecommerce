package routes

import (
	"bookecom/controllers/book"
	"bookecom/middleware"

	"github.com/gofiber/fiber/v2"
)

func BookRoutes(group fiber.Router){

	bookGroup := group.Group("/book")

	bookGroup.Post("/", middleware.TokenValidation, book.CreateBookController)
	bookGroup.Get("/allBooks", middleware.TokenValidation, book.FindAllBooks)

	bookGroup.Route("/:bookId", func(router fiber.Router) {
		router.Get("", book.FindBookByIdController)
		router.Delete("", middleware.TokenValidation, book.DeleteBook)
	})

}