package order

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type Order struct {
	ID         uuid.UUID   `json:"id"`
	PublicID   string      `json:"publicId"`
	Total      float64     `json:"total"`
	OrderItems []OrderItem `json:"items"`
	Created    string      `json:"created"`
	CreatedAt  time.Time   `json:"createdAt"`
	UpdatedAt  *time.Time  `json:"updatedAt"`
}

type OrderItem struct {
	ID        uuid.UUID `json:"id"`
	SKU       string    `json:"sku"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	Quantity  int       `json:"quantity"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
