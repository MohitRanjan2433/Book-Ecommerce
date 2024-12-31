package service

import (
	"bookecom/database"
	"bookecom/models"

	"github.com/google/uuid"
	"github.com/lib/pq"
)

func GetOrders(orderIDs pq.StringArray) ([]models.Order, error){

	var orders []models.Order

	for _, orderID := range orderIDs {
		var order models.Order

		err := database.DB.First(&order, "id = ?", orderID).Error
		if err != nil{
			return nil, err
		}

		orders = append(orders, order)
	}

	return orders, nil
}

func GetOrderById(orderID string) (*models.Order, error){

	orderuid := uuid.MustParse(orderID)
	var order models.Order
	result := database.DB.Where("id = ?", orderuid).First(&order)
	if result.Error != nil{
		return nil, result.Error
	}

	return &order, nil
}

func GetOrderByCartID(cartID uuid.UUID) ([]models.Order, error){

	var orders []models.Order
	result := database.DB.Where("cart_id = ?", cartID).Find(&orders)
	if result.Error != nil{
		return nil, result.Error
	}

	return orders, nil
}