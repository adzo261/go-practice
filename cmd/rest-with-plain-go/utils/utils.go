package utils

import (
	"encoding/json"
	"io"
)

func FromJSON(reader io.Reader, obj interface{}) error {
	e := json.NewDecoder(reader)
	return e.Decode(obj)
}

func ToJSON(writer io.Writer, obj interface{}) error {
	e := json.NewEncoder(writer)
	return e.Encode(obj)
}
