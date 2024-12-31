package book

import (
	"bookecom/models"
	bookSchema "bookecom/schemas/book"
	"bookecom/service"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)


func FindAllBooks(c *fiber.Ctx) error {

	title := c.Query("title", "")
	author := c.Query("author", "")

	var books []models.Book

	// Call the service to get books by title or author
	books, err := service.GetBookByTitleOrAuthor(title, author)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	bookResponses := bookSchema.MapBookToResponse(books)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"results": len(books),
		"books": bookResponses,
	})	
}


func FindBookByIdController(c *fiber.Ctx) error{

	bookId := uuid.MustParse(c.Params("bookId"))

	book, err := service.GetBookById(bookId)
	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"message": err.Error(),
		})
	}

	bookResponse := bookSchema.MapBookDetailToResponse(*book)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"book": bookResponse,
	})
}

func FindBookByUserId(c *fiber.Ctx) error {

	userID := c.Locals("userID").(uuid.UUID)

	user, err := service.GetUserById(userID)
	if err != nil{
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status": "error",
			"message": err.Error(),
		})
	}

	books := []models.Book{}

	for _, bookID := range user.BooksBought{
		book, err := service.GetBookById(uuid.MustParse(bookID))
		if err != nil{
			if err.Error() == "record not found"{
				continue
			}else{
				return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"status": "error",
					"message": err.Error(),
				})
			}
		}

		fmt.Println(book)
		books = append(books, *book)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status": "success",
		"result": len(books),
		"books": books,
	})
}