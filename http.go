package zonotools

import (
	"bufio"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
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

func CopyFile(orgFile multipart.File, copyFilePath string) error {
	f, err := os.OpenFile(copyFilePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		//fmt.Println(err)
		return err
	}
	defer f.Close()
	// バッファリングしながらやらないと大きいファイルはメモリがきびしい
	reader := bufio.NewReader(orgFile)

	// 書き込み処理、読み込みも書き込みもバッファしながら行う、OOM Killer 対策
	wfp, err := os.Create(copyFilePath)
	defer wfp.Close()
	if err != nil {
		//fmt.Println(err)
		return err
	}
	bufferSize := 8192
	writer := bufio.NewWriterSize(wfp, bufferSize)
	defer writer.Flush()
	bytes := make([]byte, bufferSize)
	fileSize := 0
	for {
		ret, err := reader.Read(bytes)
		if err != nil && err != io.EOF {
			//fmt.Println(err)
			//io.WriteString(w, "Failed.")
			return err
		}
		if ret == 0 {
			break
		}
		fileSize += ret
		if bufferSize != ret {
			// バッファよりも read したサイズが小さい場合、
			// 0x00 の NULL が末尾にはいってしまうため、必要なデータ部分だけを抽出して Write する
			writer.Write(bytes[:ret])
		} else {
			writer.Write(bytes)
		}
	}
	writer.Flush()
	return nil
}
