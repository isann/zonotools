package zonotools

import (
	"io/ioutil"
	"net/http"
)

func IsError(err error) bool {
	return err != nil
}
func IsNoError(err error) bool {
	return err == nil
}
func IsNotError(err error) bool {
	return IsNoError(err)
}
func StringResponseBody(resp *http.Response) string {
	b, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		return string(b)
	}
	return ""
}
func Print() {
	print("v1.0.0")
}
