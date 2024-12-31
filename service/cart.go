package service

import (
	"bookecom/database"
	"bookecom/models"

	"github.com/google/uuid"
)

func GetCartById(cartId uuid.UUID) (*models.Cart, error){

	var cart models.Cart

	result := database.DB.Find(&cart, "id = ?", cartId)
	if result.Error != nil{
		return nil, result.Error
	}

	return &cart, nil
}

func GetCartByUserId(userId uuid.UUID) (*models.Cart, error){

	var cart models.Cart
	result := database.DB.First(&cart, "user_id = ?", userId)
	if result.Error != nil{
		return nil, result.Error
	}

	return &cart, nil
}

func CreateCart(userId uuid.UUID) (*models.Cart, error){

	newCart := models.Cart{
		UserID: userId,
		CreatedAt: database.DB.NowFunc(),
		UpdatedAt: database.DB.NowFunc(),
		Items: []models.CartItem{},
		Active: true,
	}

	result := database.DB.Create(&newCart)
	if result.Error != nil{
		return nil, result.Error
	}

	return &newCart, nil
}

func DeleteUserCart(userId uuid.UUID) error{

	var cart models.Cart
	result := database.DB.Where("user_id = ?", userId).First(&cart)
	if result.Error != nil{
		return result.Error
	}

	result = database.DB.Delete(&models.CartItem{}, "cart_id = ?", cart.ID)

	result = database.DB.Delete(&cart)

	return result.Error
}