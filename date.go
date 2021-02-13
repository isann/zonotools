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

// time.Time オブジェクトを指定した日付レイアウトの文字列に変換
func FormatTime(time time.Time, layout string) string {
	return time.Format(layout)
}

// 文字列の日付レイアウトを指定して Time オブジェクトに変換
func ParseTime(timeString, layout string) (time.Time, error) {
	return time.Parse(layout, timeString)
}

const (
	// default UnitNanosecond
	UnitNanosecond = iota
	UnitMicrosecond
	UnitMillisecond
	UnitSecond
	UnitMinute
	UnitHour
)

// 日付を加減算する、unit は UnitNanosecond ~ UnitHour から指定
func AddTime(t time.Time, num int64, unit int) time.Time {
	var u time.Duration
	switch unit {
	case 0:
		u = time.Nanosecond
	case 1:
		u = time.Microsecond
	case 2:
		u = time.Millisecond
	case 3:
		u = time.Second
	case 4:
		u = time.Minute
	case 5:
		u = time.Hour
	default:
		u = time.Nanosecond
	}
	return t.Add(time.Duration(num) * u)
}
