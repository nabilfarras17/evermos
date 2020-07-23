package main

import (
	"github.com/evermos/race-condition-order/config"
	"github.com/evermos/race-condition-order/nsq/consumer"
	"github.com/evermos/race-condition-order/nsq/producer"
	"github.com/evermos/race-condition-order/pkg/handler"
	"github.com/evermos/race-condition-order/pkg/order"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

func main() {
	conf := config.InitConfig()

	// Setup producer
	producer, err := producer.NewNSQProducer(conf.ProducerHost)
	if err != nil {
		log.Fatalf("%v", err)
		return
	}

	// Setup dependency for each domain
	orderService := order.NewService(producer)
	orderHandler := order.NewOrderHandler(orderService)

	// Setup consumer
	consumer := consumer.NewNSQConsumer(
		conf.ConsumerName,
		conf.ConsumerEnabled,
		conf.ConsumerHost,
		conf.ConsumerTopic,
		orderService,
	)
	err = consumer.Start()
	if err != nil {
		log.Fatalf("%v", err)
		return
	}

	// Setup rootHandler
	router := mux.NewRouter()
	rh := handler.New(conf, router, &orderHandler)
	rh.InitRoutes()
	rh.Run()
}
