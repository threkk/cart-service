package item

import (
	"time"

	"gorm.io/gorm"
)

// Item Item GORM model
type Item struct {
	gorm.Model
	CartID    uint
	OrderID   uint
	Price     float64
	ProductID uint
	Quantity  uint
}

// Response Item JSON response
type Response struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	CartID    uint      `json:"cart_id"`
	Price     float64   `json:"price"`
	ProductID uint      `json:"product_id"`
	Quantity  uint      `json:"quantity"`
}
