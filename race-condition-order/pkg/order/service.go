package order

import (
	"context"
	"github.com/evermos/race-condition-order/nsq/producer"
)

type Service struct {
	producer producer.NSQProducer
}

func NewService(producer producer.NSQProducer) Service {
	return Service{
		producer: producer,
	}
}

func (s *Service) SaveOrder(ctx context.Context) error {
	return nil
}
