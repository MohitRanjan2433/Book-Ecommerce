package database

import (
	"bookecom/models"
	"fmt"
	"log"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	log.Println("Running Migrations")

	err := db.AutoMigrate(
		&models.User{},
		&models.RefreshToken{},
		&models.Book{},
		&models.Cart{},
		&models.CartItem{},
		&models.Order{},
		&models.Review{},
	)
	if err != nil {
		fmt.Println("Migration error")
		return
	}

	log.Println("🚀 Migrations completed")
}