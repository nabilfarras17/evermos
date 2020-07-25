package order

import (
	"context"
	"github.com/evermos/race-condition-order/config"
	"github.com/evermos/race-condition-order/nsq/producer"
	"github.com/evermos/race-condition-order/pkg/pcatalogue"
	"github.com/evermos/race-condition-order/pkg/util"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"time"
)

type Service struct {
	config            *config.Config
	producer          producer.NSQProducer
	pcatalogueService pcatalogue.Service
	repository        Repository
}

func NewService(
	conf *config.Config,
	producer producer.NSQProducer,
	pcatalogueService pcatalogue.Service,
) Service {
	orderRepository := NewRepository()
	return Service{
		config:            conf,
		producer:          producer,
		pcatalogueService: pcatalogueService,
		repository:        orderRepository,
	}
}

func (s *Service) SaveOrder(ctx context.Context, input SaveOrderRequest) (res Order, err error) {
	skus := make([]string, 0)
	for _, item := range input.Items {
		skus = append(skus, item.SKU)
	}
	products := s.pcatalogueService.BulkGetProductBySkus(ctx, skus)
	if len(products) == 0 {
		err = errors.New("[Error] SKU not found!")
		return
	}

	productMap := make(map[string]pcatalogue.Product, 0)
	for _, product := range products {
		productMap[product.SKU] = product
	}

	orderItems := make([]OrderItem, 0)
	total := float64(0)
	for _, item := range input.Items {
		if product, ok := productMap[item.SKU]; ok {
			total += product.Price
			if item.Quantity > product.Quantity {
				err = errors.Errorf("[Error] SKU: %v out of stock!", product.SKU)
				return
			}
			orderItem := OrderItem{
				ID:        uuid.NewV4(),
				SKU:       product.SKU,
				Name:      product.Name,
				Price:     product.Price,
				Quantity:  item.Quantity,
				CreatedAt: time.Now(),
			}
			orderItems = append(orderItems, orderItem)
		} else {
			err = errors.New("[Error] SKU not found!")
			return
		}
	}
	if input.Total != total {
		err = errors.New("[Error] Total input is invalid!")
		return
	}

	order := Order{
		ID:         uuid.NewV4(),
		PublicID:   input.PhoneNumber + util.RandomString(5),
		Total:      input.Total,
		Status:     PendingOrderStatus,
		Created:    input.Created,
		OrderItems: orderItems,
		CreatedAt:  time.Now(),
	}

	s.repository.InsertOrder(ctx, order)
	err = s.producer.Emit(s.config.Topic, order)
	if err != nil {
		log.Errorf("%v", err)
		return
	}
	return order, nil
}

func (s *Service) ProcessOrder(ctx context.Context, order Order) (err error) {
	log.Infof("Processing order: %v on publicID: %v", order, order.PublicID)
	skus := make([]string, 0)
	for _, orderItem := range order.OrderItems {
		skus = append(skus, orderItem.SKU)
	}
	products := s.pcatalogueService.BulkGetProductBySkus(ctx, skus)
	if len(products) == 0 {
		err = errors.New("[Error] SKU not found!")
		return
	}

	productMap := make(map[string]pcatalogue.Product, 0)
	for _, product := range products {
		productMap[product.SKU] = product
	}

	isQuantityExceed := true
	for _, item := range order.OrderItems {
		if product, ok := productMap[item.SKU]; ok {
			if item.Quantity > product.Quantity {
				log.Errorf("[Error] SKU: %v out of stock!", product.SKU)
				isQuantityExceed = false
				break
			}
			s.pcatalogueService.ReduceQuantityBySKU(ctx, item.SKU, item.Quantity)
		} else {
			err = errors.New("[Error] SKU not found!")
			return
		}
	}
	if isQuantityExceed {
		order.Status = SuccessOrderStatus
	} else {
		order.Status = CancelledOrderStatus
	}
	_, err = s.repository.UpdateOrder(ctx, order)
	if err != nil {
		return
	}
	log.Infof("Success processed order: %v on publicID: %v", order, order.PublicID)
	return nil
}

func (s *Service) GetOrderByPublicID(ctx context.Context, publicID string) (res Order, err error) {
	return s.repository.GetOrderByPublicID(ctx, publicID)
}
