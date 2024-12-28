package utils

import (
	"bookecom/config"
	"bookecom/models"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func generateToken(user *models.User, secretKey string, expiration time.Duration) (string, error){
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": user.ID,
		"exp":       time.Now().Add(expiration).Unix(),
		"role":      user.Role,
		"username":  user.Username,
		"email":     user.Email,
		"createdAt": user.CreatedAt.Unix(),
	})


	return token.SignedString([]byte(secretKey))
}

func GenerateAccessToken(user *models.User) (string, error){
	config, _ := config.LoadConfig(".")
	return generateToken(user, config.AccessTokenSecret, config.AccessTokenExpiry)
}

func GenerateRefreshToken(user *models.User) (string, error){
	config, _ := config.LoadConfig(".")
	return generateToken(user, config.RefreshTokenSecret, config.RefreshTokenExpiry)
}