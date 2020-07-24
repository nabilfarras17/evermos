package util

import (
	"encoding/json"
	"github.com/pkg/errors"
	// log "github.com/sirupsen/logrus"
)

// ConvertInterfaceUsingJSON converts interface using json
func ConvertInterfaceUsingJSON(source interface{}, destination interface{}) (err error) {
	byteResult, err := json.Marshal(source)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	err = json.Unmarshal(byteResult, destination)
	if err != nil {
		err = errors.WithStack(err)
		return
	}
	return
}
