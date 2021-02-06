package zonotools

import "encoding/json"

func DecodeJson(data interface{}) ([]byte, error) {
	return json.Marshal(data)
}

// Data required pointer
func EncodeJson(jsonString string, data interface{}) error {
	return json.Unmarshal([]byte(jsonString), data)
}
