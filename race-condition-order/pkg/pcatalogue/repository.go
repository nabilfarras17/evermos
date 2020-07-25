package pcatalogue

import (
	"context"
	"fmt"
	"github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"sync"
	"time"
)

var dummyProductMap map[string]Product
var lock = sync.RWMutex{}

type Repository struct{}

func NewRepository() Repository {
	dummyProductMap = make(map[string]Product, 0)
	generateProductCount := 2
	log.Infof("Start generate %v dummy product", generateProductCount)
	i := 0
	for i <= generateProductCount {
		product := Product{
			ID:        uuid.NewV4(),
			SKU:       fmt.Sprintf("product-%v", i+1),
			Name:      fmt.Sprintf("product-%v", i+1),
			Price:     300000,
			Quantity:  2,
			CreatedAt: time.Now(),
		}
		dummyProductMap[product.SKU] = product
		i++
	}
	log.Infof("Successfully generate %v dummy products", generateProductCount)
	return Repository{}
}

func (r *Repository) BulkGetProductBySkus(ctx context.Context, skus []string) (res []Product) {
	lock.RLock()
	defer lock.RUnlock()
	for idx := range skus {
		if product, ok := dummyProductMap[skus[idx]]; ok {
			res = append(res, product)
		} else {
			return []Product{}
		}
	}
	return res
}

func (r *Repository) ReduceQuantityBySKU(ctx context.Context, sku string, quantity int) (res Product) {
	lock.Lock()
	defer lock.Unlock()
	if product, ok := dummyProductMap[sku]; ok {
		currentQty := product.Quantity - quantity
		if currentQty >= 0 {
			product.Quantity = currentQty
			dummyProductMap[product.SKU] = product
		}
		return product
	}
	return Product{}
}

func (r *Repository) ResolveProductBySKU(ctx context.Context, sku string) *Product {
	lock.RLock()
	defer lock.RUnlock()
	if product, ok := dummyProductMap[sku]; ok {
		return &product
	}
	return nil
}
