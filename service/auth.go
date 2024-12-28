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

func LoginUser(payload *userSchema.LoginUserSchema) (models.AuthResponse, error){
	var user models.User
	result := database.DB.Where("email = ?", payload.Email).First(&user)
	if result.Error != nil{
		return models.AuthResponse{}, result.Error
	}

	fmt.Println(user)

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password))
	if err != nil{
		return models.AuthResponse{}, err
	}

	authResponse, err := GenerateAuthTokens(&user)
	if err != nil{
		return models.AuthResponse{}, err
	}

	return authResponse, nil
}

func VerifyOTP(userID uuid.UUID, otp string) error {
	var user models.User

	result := database.DB.Where("id = ? AND otp = ?", userID, otp).First(&user)
	if result.Error != nil{
		return result.Error
	}

	user.Verified = true
	user.Otp = ""
	result = database.DB.Save(&user)
	if result.Error != nil{
		return result.Error
	}

	return nil
}

func GenerateAuthTokens(user  *models.User) (models.AuthResponse, error){
	access_token, err := utils.GenerateAccessToken(user)
	if err != nil{
		return models.AuthResponse{}, err
	}

	refresh_token, err := utils.GenerateRefreshToken(user)
	if err != nil{
		return models.AuthResponse{}, err
	}

	refreshTokenEntry := models.RefreshToken{
		UserID: user.ID,
		Token: refresh_token,
	}

	result := database.DB.Create(&refreshTokenEntry)
	if result.Error != nil{
		return models.AuthResponse{}, err
	}

	authResponse := models.AuthResponse{
		UserID: user.ID,
		AccessToken: access_token,
		RefreshToken: refresh_token,
		Verified: user.Verified,
	}

	return authResponse, nil
}
