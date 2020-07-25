package pcatalogue

import (
	"github.com/satori/go.uuid"
	"time"
)

type Product struct {
	ID        uuid.UUID  `json:"id"`
	SKU       string     `json:"sku"`
	Name      string     `json:"name"`
	Price     float64    `json:"price"`
	Quantity  int        `json:"quantity"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}
