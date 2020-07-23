package main

import (
	"github.com/evermos/race-condition-order/config"
	"github.com/evermos/race-condition-order/pkg/handler"
	"github.com/evermos/race-condition-order/pkg/order"
	"github.com/gorilla/mux"
)

func main() {
	conf := config.InitConfig()
	// Setup dependency for each domain
	orderService := order.NewService()
	orderHandler := order.NewOrderHandler(orderService)

	// Setup rootHandler
	router := mux.NewRouter()
	rh := handler.New(conf, router, &orderHandler)
	rh.InitRoutes()
	rh.Run()
}
