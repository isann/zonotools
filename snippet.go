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
func StringResponseBody(resp *http.Response) string {
	b, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		return string(b)
	}
	return ""
}
