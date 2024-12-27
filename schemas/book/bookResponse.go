package book

import (
	"bookecom/models"
	"encoding/json"

	"github.com/google/uuid"
)

type BookResponse struct {
	ID          uuid.UUID `json:"id,omitempty"`
	Title       string    `json:"title,omitempty"`
	Author      string    `json:"author,omitempty"`
	Description string    `json:"description,omitempty"`
	Price       float64   `json:"price,omitempty"`
	CoverImages []string  `json:"cover_images,omitempty"`
}


func MapBookToResponse(books []models.Book) []BookResponse {

	bookResponses := make([]BookResponse, len(books))
	for i, book := range books{
		var coverImages []string
		_ = json.Unmarshal([]byte(book.CoverImages), &coverImages)

		bookResponses[i] = BookResponse{	
			ID: book.ID,
			Title: book.Title,
			Author: book.Author,
			Description: book.Description,
			Price: book.Price,
			CoverImages: coverImages,
		}
	}

	return bookResponses
}