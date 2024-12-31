package service

import (
	"bookecom/database"
	"bookecom/models"

	"github.com/google/uuid"
)

func GetAllUsers() ([]models.User, error){
	var users []models.User

	result := database.DB.Where(&users)
	if result.Error != nil{
		return nil, result.Error
	}

	return users,nil
}

func GetUserById(userID uuid.UUID) (models.User, error){

	var user models.User
	result := database.DB.Where("id = ?", userID).First(&user)
	if result.Error != nil{
		return user, result.Error
	}

	return user,nil
}

func GetUserByUserName(username string) (models.User, error) {
	var user models.User
	result := database.DB.Where("username = ?", username).First(&user)
	if result.Error != nil {
		return models.User{}, result.Error
	}
	return user, nil
}

func GetUserByEmail(email string) (models.User, error) {
	var user models.User
	result := database.DB.Where("email = ?", email).First(&user)
	if result.Error != nil{
		return user, result.Error
	}
	return user, nil
}

func GetUserByRole(role string) ([]models.User, error){
	var users []models.User

	result := database.DB.Where("role = ?", role).Find(&users)
	if result.Error != nil{
		return nil, result.Error
	}

	return users, nil
}

func UpdateUser(user models.User) (models.User, error) {

	result := database.DB.Save(&user)
	if result.Error != nil {
		return models.User{}, result.Error
	}
	return user, nil
}

