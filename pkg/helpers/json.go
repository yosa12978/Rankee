package helpers

import (
	"encoding/json"
	"io"
)

func WriteJson(w io.Writer, payload interface{}) {
	json.NewEncoder(w).Encode(payload)
}

func DecodeJson(r io.Reader, obj string) interface{} {
	var res interface{}
	json.Unmarshal([]byte(obj), &res)
	return res
}
