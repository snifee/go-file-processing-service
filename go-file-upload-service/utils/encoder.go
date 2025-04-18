package utils

import (
	"bytes"
	"encoding/json"
)

func JSONSerializer(msg any) ([]byte, error) {
	var b bytes.Buffer
	encoder := json.NewEncoder(&b)
	err := encoder.Encode(msg)
	return b.Bytes(), err
}

func JSONDeserializer[T any](b []byte) (T, error) {
	var result T
	buffer := bytes.NewBuffer(b)
	decoder := json.NewDecoder(buffer)
	err := decoder.Decode(&result)
	return result, err
}
