package service

import (
	"bookecom/database"
	"bookecom/models"
)

func GetUserByEmail(email string) (models.User, error) {
	var user models.User
	result := database.DB.Where("email = ?", email).First(&user)
	if result.Error != nil{
		return user, result.Error
	}
	return user, nil
}