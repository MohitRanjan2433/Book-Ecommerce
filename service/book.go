package service

import (
	"bookecom/database"
	"bookecom/models"
	bookSchemas "bookecom/schemas/book"
	"encoding/json"
	"strconv"
)

func CreateBook(payload *bookSchemas.CreateBookSchema) (*models.Book, error){

	price, err := strconv.ParseFloat(payload.Price, 64)
	if err != nil{
		return nil, err
	}

	quantity, err := strconv.Atoi(payload.Quantity)
	if err != nil{
		return nil, err
	}

	coverImagesJSON, err := json.Marshal(payload.CoverImages)
	if err != nil{
		return nil, err
	}

	book := &models.Book{
		ISBN: payload.ISBN,
		Title: payload.Title,
		Author: payload.Author,
		Description: payload.Description,
		Genre: payload.Genre,
		Price: price,
		Quantity: quantity,
		FullText: payload.FullText,
		Sample: payload.Sample,
		CoverImages: string(coverImagesJSON),
		UserID: payload.UserID,
	}

	result := database.DB.Create(&book)
	if result.Error != nil{
		return nil, result.Error
	}

	return book, nil
}

func GetBookByTitleOrAuthon(title string, Author string) ([]models.Book, error){

	var books []models.Book

	if title != ""{
		result := database.DB.Where("title LIKE ?", "%"+title+"%").Find(&books)	
		if result.Error != nil{
			return nil, result.Error
		}
	}else{
		result := database.DB.Where("author LIKE ?", "%"+Author+"%").Find(&books)
		if result.Error != nil{
			return nil, result.Error
		}
	}

	return books, nil
}