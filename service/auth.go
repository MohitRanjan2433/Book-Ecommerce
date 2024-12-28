package service

import (
	"bookecom/config"
	"bookecom/database"
	"bookecom/models"
	userSchema "bookecom/schemas/user"
	"bookecom/utils"
	"fmt"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)


func SignupUser(payload *userSchema.RegisterUserSchema) (models.User, error){
	config, _ := config.LoadConfig(".")

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(payload.Password), bcrypt.DefaultCost)
	if err != nil{
		return models.User{}, err
	}

	otp, _ := utils.GenerateOTP(6)
	verified := false
	role := "user"

	if config.Production != true{
		otp = "123456"
		verified = true
		role = payload.Role
	}

	newUser := models.User{
		Username:     payload.Username,
		Email:        payload.Email,
		Password:     string(hashedPassword),
		PhoneNumber:  payload.PhoneNumber,
		Verified:     verified,
		Role:         role,
		Otp:          otp,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		CartId:       uuid.Nil,
	}

	result := database.DB.Create(&newUser)
	if result.Error != nil{
		return models.User{}, result.Error
	}

	if config.Production == true{
		body := fmt.Sprintf("Dear User,\n\nWelcome to the App! Thank you for joining us.\n\n"+
			"To complete your registration, please enter the following One-Time Password (OTP):\n\n"+
			"OTP: %s\n\n"+
			"This OTP is valid for a limited time only. Please keep it confidential and do not share it with anyone.\n\n", otp,
		)

		msg := fmt.Sprintf("Books App\n%s", body)

		_, err := utils.SendEmail(payload.Email, msg)
		if err != nil{
			return models.User{}, err
		}
	}

	return newUser, nil
}
