package cart

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

var userTable = map[string]uint{
	"admin": 1,
	"user":  2,
}

// Cart Cart GORM model
type Cart struct {
	gorm.Model
	UserID uint `gorm:"notNull;index"`
}

// CartResponse Cart JSON response
type CartResponse struct {
	ID        uint           `json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"update_at"`
	UserID    uint           `json:"user_id"`
	Items     []ItemResponse `json:"items"`
}

// Item Item GORM model
type Item struct {
	gorm.Model
	CartID    uint `gorm:"index"`
	Cart      Cart
	OrderID   uint `gorm:"index"`
	Price     float64
	ProductID uint
	Quantity  uint `gorm:"default:1"`
}

// ItemResponse Item JSON response
type ItemResponse struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	CartID    uint      `json:"cart_id"`
	Price     float64   `json:"price"`
	ProductID uint      `json:"product_id"`
	Quantity  uint      `json:"quantity"`
}

// GetUserID This should be a request to another service, and it should be also
// in a different layer (likely in a service layer) but for simplicty I placed
// it here
func GetUserID(user string) (uint, error) {
	if val, ok := userTable[user]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("id not found for user %s", user)
}
