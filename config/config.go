package config

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost             string
	DBPort             string
	DBUserName         string
	DBUserPassword     string
	DBName             string
	Port               string
	Email              string
	EmailPassword      string
	Production         bool
	AccessTokenSecret  string
	RefreshTokenSecret string
	AccessTokenExpiry  time.Duration
	RefreshTokenExpiry time.Duration
}

func LoadConfig(path string) (Config, error) {
	// Load the .env file
	if err := godotenv.Load(".env"); err != nil {
		return Config{}, err
	}

    log.Printf("PRODUCTION value: %s", os.Getenv("PRODUCTION"))
	production  := os.Getenv("PRODUCTION") == "true"

	// Parse AccessTokenExpiry and RefreshTokenExpiry
	accessTokenExpiry, err := time.ParseDuration(os.Getenv("ACCESS_TOKEN_EXPIRY"))
	if err != nil {
		log.Printf("Error parsing ACCESS_TOKEN_EXPIRY: %v, defaulting to 15m", err)
		accessTokenExpiry = 15 * time.Minute
	}

	refreshTokenExpiry, err := time.ParseDuration(os.Getenv("REFRESH_TOKEN_EXPIRY"))
	if err != nil {
		log.Printf("Error parsing REFRESH_TOKEN_EXPIRY: %v, defaulting to 7d", err)
		refreshTokenExpiry = 7 * 24 * time.Hour
	}

	// Build the Config struct
	return Config{
		DBHost:             os.Getenv("DB_HOST"),
		DBPort:             os.Getenv("DB_PORT"),
		DBUserName:         os.Getenv("DB_USER_NAME"),
		DBUserPassword:     os.Getenv("DB_USER_PASSWORD"),
		DBName:             os.Getenv("DB_NAME"),
		Port:               os.Getenv("PORT"),
		Email:              os.Getenv("EMAIL"),
		EmailPassword:      os.Getenv("EMAIL_PASSWORD"),
		Production:         production,
		AccessTokenSecret:  os.Getenv("ACCESS_TOKEN_SECRET"),
		RefreshTokenSecret: os.Getenv("REFRESH_TOKEN_SECRET"),
		AccessTokenExpiry:  accessTokenExpiry,
		RefreshTokenExpiry: refreshTokenExpiry,
	}, nil
}
