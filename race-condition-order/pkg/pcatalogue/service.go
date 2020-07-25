package pcatalogue

import (
	"context"
	"github.com/pkg/errors"
)

type Service struct {
	repository Repository
}

func NewService() Service {
	pcatRepository := NewRepository()
	return Service{
		repository: pcatRepository,
	}
}

func (s *Service) BulkGetProductBySkus(ctx context.Context, skus []string) (res []Product) {
	return s.repository.BulkGetProductBySkus(ctx, skus)
}

func (s *Service) ReduceQuantityBySKU(ctx context.Context, sku string, quantity int) (res Product) {
	return s.repository.ReduceQuantityBySKU(ctx, sku, quantity)
}

func (s *Service) ResolveProductBySKU(ctx context.Context, sku string) (res Product, err error) {
	product := s.repository.ResolveProductBySKU(ctx, sku)
	if product == nil {
		return Product{}, errors.Errorf("Product with sku: %v not found!", sku)
	}
	return *product, nil
}
