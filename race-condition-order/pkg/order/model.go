package order

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

var (
	PendingOrderStatus   = "pending"
	CancelledOrderStatus = "cancelled"
	SuccessOrderStatus   = "success"
)

type Order struct {
	ID         uuid.UUID   `json:"id"`
	PublicID   string      `json:"publicId"`
	Total      float64     `json:"total"`
	OrderItems []OrderItem `json:"items"`
	Created    string      `json:"created"`
	Status     string      `json:"status"`
	CreatedAt  time.Time   `json:"createdAt"`
	UpdatedAt  *time.Time  `json:"updatedAt"`
}

type OrderItem struct {
	ID        uuid.UUID  `json:"id"`
	SKU       string     `json:"sku"`
	Name      string     `json:"name"`
	Price     float64    `json:"price"`
	Quantity  int        `json:"quantity"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}
