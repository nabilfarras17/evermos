package request

import (
	"encoding/json"
	"io"
)

func ReadRequest(body io.ReadCloser, payload interface{}) (err error) {
	err = json.NewDecoder(body).Decode(&payload)
	if err != nil {
		return
	}
	return
}
