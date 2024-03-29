package zonotools

import (
	"io/ioutil"
	"log"
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
func IsExistKey[T comparable, U any](m map[T]U, key T) (U, bool) {
	if val, ok := m[key]; ok {
		return val, ok
	} else {
		return val, ok
	}
}

// IsExistsKey は、マップにキーが存在するかどうかを判定します。
// m が map ではない場合も false になります。
//
// Deprecated: Generics によりこの関数での reflect による処理は不要
func IsExistsKey[T comparable, U any](m map[T]U, k T) bool {
	defer func() {
		_ = recover()
	}()
	mapValue := reflect.ValueOf(m)
	keyValue := reflect.ValueOf(k)
	v := mapValue.MapIndex(keyValue)
	return v != reflect.Value{}
}

// Mapper は、 構造体 org のプロパティを copy にマッピングします。
// 構造体は同じ型でなくても受け渡しでき、プロパティ名が同じものをコピーします。
//
// 第一引数の構造体の値を第二引数の構造体にコピーします。
// コピーするフィールド名は同じフィールド名のもののみです。
func Mapper[T, U any](org *T, copy *U) {

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

// ToPtrTime は、文字列を time.Time 型に変換します。
//
// layout には日付文字列のレイアウト、 timeString には日付文字列、 local にはタイムゾーンを指定します。
func ToPtrTime(layout, timeString string, local *time.Location) *time.Time {
	if layout == "" {
		layout = "2006/01/02 15:04.05"
	}
	nowTime, _ := time.ParseInLocation(layout, timeString, local)
	return &nowTime
}

// Timer は、処理時間を計測します。
type Timer struct {
	startTime int64
	endTime   int64
	// nano second
	elapsed int64
}

func (timer *Timer) New() *Timer {
	return &Timer{}
}

func (timer *Timer) NewAndStart() *Timer {
	t := &Timer{}
	t.Start()
	return t
}

func (timer *Timer) Start() {
	timer.startTime = time.Now().UnixNano()
}

func (timer *Timer) End() {
	timer.endTime = time.Now().UnixNano()
	timer.elapsed = timer.endTime - timer.startTime
}

func (timer *Timer) PrintElapsed(vars ...interface{}) {
	e := timer.elapsed / int64(time.Millisecond)
	log.Println("elapsed[ms]:", e, vars)
}

// Elapsed は、処理時間を返します。
// 単位は、ナノ秒です。
func (timer *Timer) Elapsed() int64 {
	return timer.elapsed
}
