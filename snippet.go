package zonotools

import (
	"io/ioutil"
	"net/http"
	"reflect"
	"time"
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
	print("v2.0.1")
}

// IsExistKey は、マップに key が存在するかどうかを判定します。
func IsExistKey(m map[interface{}]interface{}, key interface{}) (interface{}, bool) {
	if val, ok := m[key]; ok {
		return val, ok
	} else {
		return val, ok
	}
}

// Mapper は、 構造体 org のプロパティを copy にマッピングします。
//
// 引数は構造体のアドレスを指定して、参照渡しとしてください。
// 第一引数の構造体の値を第二引数の構造体にコピーします。
// コピーするフィールド名は同じフィールド名のもののみです。
func Mapper(org, copy interface{}) {

	// check pointer
	if reflect.TypeOf(org).String()[:1] != "*" || reflect.TypeOf(copy).String()[:1] != "*" {
		// error
		return
	}
	// TODO: check struct

	// original
	rv := reflect.ValueOf(org).Elem()
	rt := rv.Type()

	// copy
	rrv := reflect.ValueOf(copy).Elem()
	rrt := rrv.Type()

	for i := 0; i < rt.NumField(); i++ {
		// フィールドの取得
		f := rt.Field(i)
		// コピー元のフィールド名とコピー先のフィールド名が同じとき、値を移し替える
		if f, ok := rrt.FieldByName(f.Name); ok {
			// org と同名フィールドの copy でのインデックス位置を取得
			barIndex := f.Index[len(f.Index)-1]
			// org の該当フィールド情報取得
			fieldInterface := reflect.ValueOf(org).Elem().Field(barIndex).Interface()
			// org の該当フィールド値取得
			fooValue := reflect.ValueOf(fieldInterface)
			// copy の該当フィールドに copy の値を設定する
			reflect.ValueOf(copy).Elem().Field(barIndex).Set(fooValue)
		}
	}
}

// ToPtrS は、文字列を参照型で返します
func ToPtrS(s string) *string {
	return &s
}

// ToPtrI64 は、 Int64 を参照型で返します
func ToPtrI64(i int64) *int64 {
	return &i
}

// ToPtrTime は、文字列を time.Time 型に変換します
//
// layout 日付文字列のレイアウトを指定します
// timeString 日付文字列
func ToPtrTime(layout, timeString string) *time.Time {
	if layout == "" {
		layout = "2006/01/02 15:04.05"
	}
	nowTime, _ := time.Parse(layout, timeString)
	// JST から UTC に
	nowTime = nowTime.Add(-9 * time.Hour)
	return &nowTime
}
