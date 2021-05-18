package helper

import "encoding/json"

func StructJSON(data interface{}) (string, error) {
	// Marshal the struct to JSON and check for errors
	b, err := json.Marshal(data)
	if err != nil {
		return string(b), err
	}

	return string(b), nil
}
