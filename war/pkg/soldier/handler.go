package soldier

import (
	"github.com/evermos/war/pkg/request"
	"github.com/evermos/war/pkg/response"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type SoldierHandler struct {
	service Service
}

func NewSoldierHandler(service Service) SoldierHandler {
	return SoldierHandler{
		service: service,
	}
}

func (sh *SoldierHandler) CreateSoldierHandler(w http.ResponseWriter, r *http.Request) {
	log.Info("MASUK SINI")
	var data CreateSoldierRequest
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
	soldier, err := sh.service.CreateSoldier(ctx, data)
	if err != nil {
		log.Errorf("%v", err)
		response.Failed(w, http.StatusInternalServerError, err)
		return
	}
	response.Success(w, http.StatusAccepted, soldier)
}
