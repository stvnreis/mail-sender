package mapper

import (
	"encoding/json"
)

func FromJson(jsonData []byte, v interface{}) error {
	err := json.Unmarshal(jsonData, v)
	if err != nil {
		return err
	}
	return nil
}

func ToJson(v interface{}) ([]byte, error) {
	jsonData, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}