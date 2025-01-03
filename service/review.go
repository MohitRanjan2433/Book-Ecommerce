package service

import (
	"bookecom/database"
	"bookecom/models"
	reviewSchema "bookecom/schemas/review"

	"github.com/google/uuid"
)

func CreateReview(userID uuid.UUID, bookID uuid.UUID, userName string, payload reviewSchema.CreateReviewSchema) (models.Review, error){
	reviews := models.Review{
		UserID: userID,
		BookID: bookID,
		Username: userName,
		Rating: float64(payload.Rating),
		Comment: payload.Comment,
	}

	err := database.DB.Create(&reviews)
	if err != nil{
		return models.Review{}, err.Error
	}

	return reviews, nil
}

func GetUserReviewByUserIDAndBookID(userID, bookID uuid.UUID) (models.Review, error){

	var review models.Review

	result := database.DB.Where("user_id = ? AND book_id = ?", userID, bookID).First(&review)
	if result.Error != nil{
		return models.Review{}, result.Error
	}

	return review, nil
}

func UpdateReview(userID uuid.UUID, bookID uuid.UUID, userName string, payload reviewSchema.CreateReviewSchema) (models.Review, error){
	
	reviews := models.Review{
		UserID: userID,
		BookID: bookID,
		Username: userName,
		Rating: float64(payload.Rating),
		Comment: payload.Comment,
	}

	result := database.DB.Model(&reviews).Where("user_id = ? AND book_id = ?", userID, bookID).Updates(reviews)
	if result.Error != nil{
		return models.Review{}, result.Error
	}

	return reviews, nil
}

func GetReviewById(reviewId uuid.UUID) (models.Review, error){

	var review models.Review
	result := database.DB.First(&review, "id = ?", reviewId)
	if result.Error != nil{
		return models.Review{}, result.Error
	}

	return review, nil
}

func GetReviewByBookId(bookId uuid.UUID)([]models.Review, error){

	var review []models.Review
	result := database.DB.Find(&review, "book_id = ?", bookId)
	if result.Error != nil{
		return []models.Review{}, result.Error
	}

	return review, nil
}

func DeleteReview(reviewId uuid.UUID) error {
	review, err := GetReviewById(reviewId)
	if err != nil{
		return err
	}

	result := database.DB.Delete(&review)
	if result.Error != nil{
		return result.Error
	}

	return nil
}