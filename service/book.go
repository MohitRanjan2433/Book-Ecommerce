package service

import (
	"bookecom/database"
	"bookecom/models"
	bookSchemas "bookecom/schemas/book"
	"encoding/json"
	"errors"
	"strconv"

	"github.com/google/uuid"
	"gorm.io/gorm"
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

func GetBookByTitleOrAuthor(title string, Author string) ([]models.Book, error){

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

func GetBookById(bookID uuid.UUID) (*models.Book, error){

	var book models.Book

	if err := database.DB.First(&book, "id = ?", bookID).Error; err != nil{
		if err == gorm.ErrRecordNotFound{
			return nil, err
		}
		return nil, err
	}

	return &book, nil
}

func GetBookByIsbn(isbn string) (*models.Book, error){

	var book models.Book

	result := database.DB.Where("isbn = ?", isbn).First(&book)
	if result.Error != nil{
		return nil, result.Error
	}

	return &book, nil
}

func UpdateBook(book *models.Book) (*models.Book, error){

	result := database.DB.Save(&book)
	if result.Error != nil{
		return nil, result.Error
	}

	return book, nil
}

func DeleteBook(userID uuid.UUID, bookID string) error{

	var book models.Book

	result := database.DB.First(&book, "id = ?", bookID)
	if result.Error != nil{
		return result.Error
	}

	if userID != book.UserID{
		return errors.New("you are not the owner of this book")
	}

	result = database.DB.Delete(&book)
	if result.Error != nil{
		return result.Error
	}

	return nil
}

