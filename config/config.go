package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost        string
    DBPort        string
    DBUserName    string
    DBUserPassword string
    DBName        string
    Port          string
}

func LoadConfig(path string) (Config, error) {
    if err := godotenv.Load(path + "/.env"); err != nil {
        log.Fatal("Error loading .env file")
    }

    return Config{
        DBHost:        os.Getenv("DB_HOST"),
        DBPort:        os.Getenv("DB_PORT"),
        DBUserName:    os.Getenv("DB_USER_NAME"),
        DBUserPassword: os.Getenv("DB_USER_PASSWORD"),
        DBName:        os.Getenv("DB_NAME"),
        Port:          os.Getenv("PORT"),
    }, nil
}