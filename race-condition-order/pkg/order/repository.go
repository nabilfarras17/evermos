package order

import (
	"context"
	"github.com/pkg/errors"
	"sync"
)

var orderMap map[string]Order
var lock = sync.RWMutex{}

type Repository struct{}

func NewRepository() Repository {
	orderMap = make(map[string]Order, 0)
	return Repository{}
}

func (r *Repository) InsertOrder(ctx context.Context, order Order) Order {
	lock.Lock()
	defer lock.Unlock()
	orderMap[order.PublicID] = order
	return order
}

func (r *Repository) UpdateOrder(ctx context.Context, order Order) (res Order, err error) {
	lock.Lock()
	defer lock.Unlock()
	if _, ok := orderMap[order.PublicID]; ok {
		orderMap[order.PublicID] = order
		res = order
		return
	}
	err = errors.Errorf("[Error] updateOrder reason: orderPublicID: %v not found!", order.PublicID)
	return
}

func (r *Repository) GetOrderByPublicID(ctx context.Context, publicID string) (res Order, err error) {
	lock.RLock()
	defer lock.RUnlock()
	if order, ok := orderMap[publicID]; ok {
		return order, nil
	} else {
		return Order{}, errors.Errorf("[Error] OrderPublicID: %v not found!", publicID)
	}
}
