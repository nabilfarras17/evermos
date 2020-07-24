package order

import (
	"context"
	"github.com/pkg/errors"
)

var orderMap map[string]Order

type Repository struct{}

func NewRepository() Repository {
	orderMap = make(map[string]Order, 0)
	return Repository{}
}

func (r *Repository) InsertOrder(ctx context.Context, order Order) Order {
	orderMap[order.PublicID] = order
	return order
}

func (r *Repository) GetOrderByPublicID(ctx context.Context, publicID string) (res Order, err error) {
	if order, ok := orderMap[publicID]; ok {
		return order, nil
	} else {
		return Order{}, errors.Errorf("OrderPublicID: %v not found!", publicID)
	}
}
