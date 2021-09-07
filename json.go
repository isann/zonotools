package zonotools

import "encoding/json"

// DecodeJson は、 第一引数の interface 型を JSON 文字列の byte 配列に変換します。
// data はポインター型でなければいけないです。
func DecodeJson(data interface{}) ([]byte, error) {
	return json.Marshal(data)
}

// EncodeJson は、 第一引数の JSON 文字列を第二引数の interface 型に変換します。
// data はポインター型でなければいけないです。
func EncodeJson(jsonString string, data interface{}) error {
	return json.Unmarshal([]byte(jsonString), data)
}
