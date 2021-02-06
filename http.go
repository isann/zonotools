package zonotools

import (
	"io/ioutil"
	"mime/multipart"
	"net/http"
)

func GetQueryParam(r *http.Request, key string) string {
	return r.URL.Query().Get(key)
}

func GetFormValue(r *http.Request, key string) string {
	return r.FormValue(key)
}

func GetFormFile(r *http.Request, key string) (multipart.File, *multipart.FileHeader, error) {
	return r.FormFile(key)
}

func GetRequestBody(r *http.Request) ([]byte, error) {
	return ioutil.ReadAll(r.Body)
}

//import 	"goji.io/pat"
//// URL /file/:fid
//// 	mux.HandleFunc(pat.Post("/file/:fid"), withVars(withData(db, getFile)))
//func getFile(w http.ResponseWriter, r *http.Request) {
//	// parse key
//	fid := pat.Param(r, "fid")
