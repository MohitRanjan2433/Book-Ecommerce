package routes

import (
	"github.com/gofiber/fiber/v2"
	"bookecom/controllers/book"
)

func BookRoutes(group fiber.Router){

	bookGroup := group.Group("/book")

	bookGroup.Post("/", book.CreateBookController)
	bookGroup.Get("/allBooks", book.FindAllBooks)

}