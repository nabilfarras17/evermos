package consumer

import (
	"github.com/evermos/race-condition-order/pkg/order"
	"github.com/nsqio/go-nsq"
	log "github.com/sirupsen/logrus"
)

type NSQConsumer struct {
	name           string
	enabled        bool
	host           string
	topic          string
	eventProcessor EventProcessor
	nsqConsumer    *nsq.Consumer
}

func NewNSQConsumer(
	name string,
	enabled bool,
	host string,
	topic string,
	orderService order.Service,
) NSQConsumer {
	eventProcessor := EventProcessor{
		orderService: orderService,
	}
	return NSQConsumer{
		name:           name,
		enabled:        enabled,
		host:           host,
		topic:          topic,
		eventProcessor: eventProcessor,
	}
}

func (nc *NSQConsumer) Name() string {
	return nc.name
}

func (nc *NSQConsumer) Enabled() bool {
	return nc.enabled
}

func (nc *NSQConsumer) Start() error {
	err := nc.createConsumer()
	if err != nil {
		return err
	}

	nc.nsqConsumer.AddHandler(&nc.eventProcessor)
	err = nc.nsqConsumer.ConnectToNSQD(nc.host)
	if err != nil {
		log.Errorf("%v", err)
		return err
	}
	return nil
}

func (nc *NSQConsumer) createConsumer() error {
	config := nsq.NewConfig()
	consumer, err := nsq.NewConsumer(nc.topic, nc.name, config)
	if err != nil {
		log.Fatal(err)
		return err
	}
	nc.nsqConsumer = consumer
	return nil
}
