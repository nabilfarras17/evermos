package consumer

import (
	"github.com/evermos/race-condition-order/pkg/order"
	"github.com/nsqio/go-nsq"
	log "github.com/sirupsen/logrus"
)

type EventProcessor struct {
	orderService order.Service
}

func NewEventProcessor(orderService order.Service) EventProcessor {
	return EventProcessor{
		orderService: orderService,
	}
}

func (e *EventProcessor) HandleMessage(m *nsq.Message) error {
	if len(m.Body) == 0 {
		return nil
	}
	log.Infof("Receive a message: %v", string(m.Body))
	return nil
}
