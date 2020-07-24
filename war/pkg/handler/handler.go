package handler

import (
	"github.com/evermos/war/config"
	"github.com/evermos/war/pkg/soldier"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type RootHandler struct {
	config         *config.Config
	router         *mux.Router
	soldierHandler *soldier.SoldierHandler
	isShuttingDown bool
}

// Initiate RootHandler
func New(config *config.Config, router *mux.Router, soldierHandler *soldier.SoldierHandler) *RootHandler {
	h := &RootHandler{
		config:         config,
		router:         router,
		soldierHandler: soldierHandler,
		isShuttingDown: false,
	}
	return h
}

func (h *RootHandler) Run() {
	// register graceful shutdown
	h.handleShutdown()
	// start server
	log.Infof("Server started at :" + h.config.Port)
	log.Fatal("", http.ListenAndServe(":"+h.config.Port, h.router))
}

// handleShutdown handlers shutdown gracefully.
func (h *RootHandler) handleShutdown() {
	ch := make(chan os.Signal, 2)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	go func(ch chan os.Signal) {
		<-ch
		defer os.Exit(0)
		duration := h.config.ShutdownPeriod
		log.Infof("Signal termination received. Waiting %v seconds to shutdown", duration.Seconds())
		h.isShuttingDown = true
		time.Sleep(duration)
		log.Info("Cleaning up resources...")
		log.Info("Bye")
	}(ch)
}
