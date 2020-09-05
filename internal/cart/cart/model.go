package cart

import (
	"time"

	"github.com/threkk/cart-service/internal/cart/item"
	"gorm.io/gorm"
)

// Cart Cart GORM model
type Cart struct {
	gorm.Model
	UserID uint
}

// Response Cart JSON response
type Response struct {
	ID        uint            `json:"id"`
	CreatedAt time.Time       `json:"created_at"`
	UpdateAt  time.Time       `json:"update_at"`
	UserID    uint            `json:"user_id"`
	Items     []item.Response `json:"items"`
}
