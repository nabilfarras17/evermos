package main

import (
	"github.com/evermos/war/config"
	"github.com/evermos/war/pkg/handler"
	"github.com/evermos/war/pkg/soldier"
	"github.com/evermos/war/pkg/weapon"
	"github.com/gorilla/mux"
)

func main() {
	conf := config.InitConfig()
	// Setup dependency for each domain
	weaponService := weapon.NewService()
	soldierService := soldier.NewService(weaponService)
	soldierHandler := soldier.NewSoldierHandler(soldierService)

	// Setup rootHandler
	router := mux.NewRouter()
	rh := handler.New(conf, router, &soldierHandler)
	rh.InitRoutes()
	rh.Run()
}
