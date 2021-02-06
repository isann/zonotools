package zonotools

import "encoding/json"

func DecodeJson(data map[interface{}]interface{}) ([]byte, error) {
	return json.Marshal(data)
}

func EncodeJson(jsonString string, data *map[interface{}]interface{}) error {
	return json.Unmarshal([]byte(jsonString), data)
}
