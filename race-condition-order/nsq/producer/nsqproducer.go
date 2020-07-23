package producer

import (
	"encoding/json"
	"github.com/nsqio/go-nsq"
	log "github.com/sirupsen/logrus"
)

type NSQProducer struct {
	producer *nsq.Producer
}

func NewNSQProducer(host string) (NSQProducer, error) {
	config := nsq.NewConfig()
	producer, err := nsq.NewProducer(host, config)
	if err != nil {
		log.Fatal(err)
		return NSQProducer{}, err
	}
	return NSQProducer{
		producer: producer,
	}, nil
}

func (np *NSQProducer) Emit(topic string, message interface{}) error {
	payload, err := json.Marshal(message)
	if err != nil {
		return err
	}

	err = np.producer.Publish(topic, payload)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
