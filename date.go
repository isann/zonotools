package zonotools

import "time"

const (
	TimeStdLongMonth      = "January"
	TimeStdMonth          = "Jan"
	TimeStdNumMonth       = "1"
	TimeStdZeroMonth      = "01"
	TimeStdLongWeekDay    = "Monday"
	TimeStdWeekDay        = "Mon"
	TimeStdDay            = "2"
	TimeStdUnderDay       = "_2"
	TimeStdZeroDay        = "02"
	TimeStdHour           = "15"
	TimeStdHour12         = "3"
	TimeStdZeroHour12     = "03"
	TimeStdMinute         = "4"
	TimeStdZeroMinute     = "04"
	TimeStdSecond         = "5"
	TimeStdZeroSecond     = "05"
	TimeStdLongYear       = "2006"
	TimeStdYear           = "06"
	TimeStdPM             = "PM"
	TimeStdpm             = "pm"
	TimeStdTZ             = "MST"
	TimeStdISO8601TZ      = "Z0700"  // prints Z for UTC
	TimeStdISO8601ColonTZ = "Z07:00" // prints Z for UTC
	TimeStdNumTZ          = "-0700"  // always numeric
	TimeStdNumShortTZ     = "-07"    // always numeric
	TimeStdNumColonTZ     = "-07:00" // always numeric
)

// FormatTime は、 time.Time オブジェクトを指定した日付レイアウトの文字列に変換します。
func FormatTime(time time.Time, layout string) string {
	return time.Format(layout)
}

// ParseTime は文字列の日付レイアウトを指定して Time オブジェクトに変換に変換します。
// タイムゾーンは UTC になります。
func ParseTime(timeString, layout string) (time.Time, error) {
	return time.Parse(layout, timeString)
}

// ParseTimeInLocation は文字列の日付レイアウトを指定して Time オブジェクトに変換に変換します。
// local でタイムゾーンを指定できます。
func ParseTimeInLocation(timeString, layout string, local *time.Location) (time.Time, error) {
	return time.ParseInLocation(layout, timeString, local)
}

const (
	// UnitNanosecond がデフォルトです
	UnitNanosecond = iota
	UnitMicrosecond
	UnitMillisecond
	UnitSecond
	UnitMinute
	UnitHour
)

// AddTime は、日付を加減算します。
// unit は UnitNanosecond ~ UnitHour から指定します。
func AddTime(t time.Time, num int64, unit int) time.Time {
	var u time.Duration
	switch unit {
	case UnitNanosecond:
		u = time.Nanosecond
	case UnitMicrosecond:
		u = time.Microsecond
	case UnitMillisecond:
		u = time.Millisecond
	case UnitSecond:
		u = time.Second
	case UnitMinute:
		u = time.Minute
	case UnitHour:
		u = time.Hour
	default:
		u = time.Nanosecond
	}
	return t.Add(time.Duration(num) * u)
}

// CompareTime は、 time.Time 型のパラメータを expr の式で比較した結果を返します。
func CompareTime(a time.Time, expr string, b time.Time) bool {
	result := false
	switch expr {
	case "==":
		result = a.Equal(b)
	case "<":
		result = a.Before(b)
	case ">":
		result = a.After(b)
	case "<=":
		result = a.Equal(b) || a.Before(b)
	case ">=":
		result = a.Equal(b) || a.After(b)
	default:
		result = false
	}
	return result
}
