package response

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func Success(w http.ResponseWriter, status int, data interface{}) {
	resp := map[string]interface{}{
		"data":  data,
		"error": nil,
	}
	js, err := json.Marshal(resp)
	if err != nil {
		resp := map[string]interface{}{
			"data":  nil,
			"error": []string{fmt.Sprintf("%s", err)},
		}
		js, _ = json.Marshal(resp)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)
}

func Failed(w http.ResponseWriter, status int, err error) {
	if status/1e2 == 4 {
		log.Warningf("%v", err)
	} else {
		log.Errorf("%v", err)
	}
	var errResp map[string]interface{}
	if err != nil {
		errCode := status
		errMsg := err.Error()
		errResp = map[string]interface{}{
			"code":    errCode,
			"message": errMsg,
		}
	}
	resp := map[string]interface{}{
		"data":  nil,
		"error": errResp,
	}
	js, _ := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)
}
