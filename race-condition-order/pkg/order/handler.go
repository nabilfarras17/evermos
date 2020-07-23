package order

import (
	"github.com/evermos/race-condition-order/pkg/request"
	"github.com/evermos/race-condition-order/pkg/response"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type OrderHandler struct {
	service Service
}

func NewOrderHandler(service Service) OrderHandler {
	return OrderHandler{
		service: service,
	}
}

func (oh OrderHandler) SaveOrder(w http.ResponseWriter, r *http.Request) {
	var data SaveOrderRequest
	err := request.ReadRequest(r.Body, &data)
	if err != nil {
		log.Errorf("Error reading input: %v", err)
		response.Failed(w, http.StatusBadRequest, err)
		return
	}

	ctx := r.Context()
	err = oh.service.SaveOrder(ctx)
	if err != nil {
		log.Errorf("%v", err)
		response.Failed(w, http.StatusInternalServerError, err)
	}
	response.Success(w, http.StatusAccepted, "OK")
}
