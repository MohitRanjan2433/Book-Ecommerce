package order

import "github.com/google/uuid"

type CreateOrderSchema struct {
	ID              uuid.UUID `json:"id,omitempty"`
	CartID          uuid.UUID `json:"cart_id,omitempty"`
}