package request

import (
	"encoding/json"
	"io"
)

func ReadRequest(body io.ReadCloser, payload interface{}) (err error) {
	defer body.Close()
	if err = json.NewDecoder(body).Decode(&payload); err != nil {
		return
	}
	return nil
}
