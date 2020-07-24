package pcatalogue

import (
	"github.com/evermos/race-condition-order/pkg/response"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type PcatalogueHandler struct {
	service Service
}

func NewPcatalogueHandler(service Service) PcatalogueHandler {
	return PcatalogueHandler{
		service: service,
	}
}

func (ph PcatalogueHandler) ResolveProductBySKU(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var sku string
	if value, ok := vars["sku"]; !ok {
		response.Failed(w, http.StatusBadRequest, errors.New("Bad Request sku!"))
		return
	} else {
		sku = value
	}

	ctx := r.Context()
	product, err := ph.service.ResolveProductBySKU(ctx, sku)
	if err != nil {
		log.Errorf("%v", err)
		response.Failed(w, http.StatusInternalServerError, err)
		return
	}
	response.Success(w, http.StatusAccepted, product)
}
