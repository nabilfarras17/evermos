package order

import (
	"github.com/evermos/race-condition-order/pkg/request"
	"github.com/evermos/race-condition-order/pkg/response"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
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

	if validErrs := data.validate(); len(validErrs) > 0 {
		err = errors.Errorf("validationError: %v", validErrs)
		response.Failed(w, http.StatusBadRequest, err)
		return
	}

	ctx := r.Context()
	res, err := oh.service.SaveOrder(ctx, data)
	if err != nil {
		log.Errorf("%v", err)
		response.Failed(w, http.StatusInternalServerError, err)
		return
	}
	response.Success(w, http.StatusAccepted, res)
}

func (oh OrderHandler) GetOrderByPublicID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var publicID string
	if value, ok := vars["publicID"]; !ok {
		response.Failed(w, http.StatusBadRequest, errors.New("Bad Request publicID!"))
		return
	} else {
		publicID = value
	}

	ctx := r.Context()
	res, err := oh.service.GetOrderByPublicID(ctx, publicID)
	if err != nil {
		log.Errorf("%v", err)
		response.Failed(w, http.StatusInternalServerError, err)
		return
	}
	response.Success(w, http.StatusAccepted, res)
}
