package zonotools

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddTime(t *testing.T) {

	layout := TimeStdLongYear + "/" + TimeStdZeroMonth + "/" + TimeStdZeroDay + " " + TimeStdHour + ":" + TimeStdZeroMinute + "." + TimeStdZeroSecond

	t.Run("normal", func(t *testing.T) {
		timeString := "2000/01/01 10:00.00"
		timeObj, _ := ParseTime(timeString, layout)
		addTime := AddTime(timeObj, 1, UnitHour)
		formatTime := FormatTime(addTime, layout)
		assert.Equal(t, "2000/01/01 11:00.00", formatTime)
	})
	t.Run("normal", func(t *testing.T) {
		timeString := "2000/01/01 10:00.00"
		timeObj, _ := ParseTime(timeString, layout)
		addTime := AddTime(timeObj, 24, UnitHour)
		formatTime := FormatTime(addTime, layout)
		assert.Equal(t, "2000/01/02 10:00.00", formatTime)
	})
	t.Run("normal", func(t *testing.T) {
		timeString := "2000/01/01 10:00.00"
		timeObj, _ := ParseTime(timeString, layout)
		addTime := AddTime(timeObj, 30, UnitMinute)
		formatTime := FormatTime(addTime, layout)
		assert.Equal(t, "2000/01/01 10:30.00", formatTime)
	})
	t.Run("normal", func(t *testing.T) {
		timeString := "2000/01/01 10:00.00"
		timeObj, _ := ParseTime(timeString, layout)
		addTime := AddTime(timeObj, 40, UnitSecond)
		formatTime := FormatTime(addTime, layout)
		assert.Equal(t, "2000/01/01 10:00.40", formatTime)
	})
	t.Run("normal", func(t *testing.T) {
		timeString := "2000/01/01 10:00.00"
		timeObj, _ := ParseTime(timeString, layout)
		addTime := AddTime(timeObj, 24*31, UnitHour)
		formatTime := FormatTime(addTime, layout)
		assert.Equal(t, "2000/02/01 10:00.00", formatTime)
	})
	t.Run("normal", func(t *testing.T) {
		timeString := "2000/01/01 10:00.00"
		timeObj, _ := ParseTime(timeString, layout)
		addTime := AddTime(timeObj, 24*(31+27), UnitHour)
		formatTime := FormatTime(addTime, layout)
		assert.Equal(t, "2000/02/28 10:00.00", formatTime)
	})
	t.Run("normal", func(t *testing.T) {
		timeString := "2000/01/01 10:00.00"
		timeObj, _ := ParseTime(timeString, layout)
		addTime := AddTime(timeObj, 24*(31+28), UnitHour)
		formatTime := FormatTime(addTime, layout)
		assert.Equal(t, "2000/02/29 10:00.00", formatTime)
	})
	t.Run("normal", func(t *testing.T) {
		timeString := "2000/01/01 10:00.00"
		timeObj, _ := ParseTime(timeString, layout)
		addTime := AddTime(timeObj, 24*(31+28+1), UnitHour)
		formatTime := FormatTime(addTime, layout)
		assert.Equal(t, "2000/03/01 10:00.00", formatTime)
	})
	t.Run("normal", func(t *testing.T) {
		timeString := "2001/01/01 10:00.00"
		timeObj, _ := ParseTime(timeString, layout)
		addTime := AddTime(timeObj, 24*(31+27), UnitHour)
		formatTime := FormatTime(addTime, layout)
		assert.Equal(t, "2001/02/28 10:00.00", formatTime)
	})
	t.Run("normal", func(t *testing.T) {
		timeString := "2001/01/01 10:00.00"
		timeObj, _ := ParseTime(timeString, layout)
		addTime := AddTime(timeObj, 24*(31+28), UnitHour)
		formatTime := FormatTime(addTime, layout)
		assert.Equal(t, "2001/03/01 10:00.00", formatTime)
	})
}

func TestFormatTime(t *testing.T) {
	//type args struct {
	//	time   time.Time
	//	layout string
	//}
	//tests := []struct {
	//	name string
	//	args args
	//	want string
	//}{
	//	// TODO: Add test cases.
	//}
	//for _, tt := range tests {
	//	t.Run(tt.name, func(t *testing.T) {
	//		if got := FormatTime(tt.args.time, tt.args.layout); got != tt.want {
	//			t.Errorf("FormatTime() = %v, want %v", got, tt.want)
	//		}
	//	})
	//}
	t.Run("", func(t *testing.T) {

	})
}

func TestParseTime(t *testing.T) {
	//type args struct {
	//	timeString string
	//	layout     string
	//}
	//tests := []struct {
	//	name    string
	//	args    args
	//	want    time.Time
	//	wantErr bool
	//}{
	//	// TODO: Add test cases.
	//}
	//for _, tt := range tests {
	//	t.Run(tt.name, func(t *testing.T) {
	//		got, err := ParseTime(tt.args.timeString, tt.args.layout)
	//		if (err != nil) != tt.wantErr {
	//			t.Errorf("PartTime() error = %v, wantErr %v", err, tt.wantErr)
	//			return
	//		}
	//		if !reflect.DeepEqual(got, tt.want) {
	//			t.Errorf("PartTime() got = %v, want %v", got, tt.want)
	//		}
	//	})
	//}
}

func TestCompareTime(t *testing.T) {

	layout := TimeStdLongYear + "/" + TimeStdZeroMonth + "/" + TimeStdZeroDay + " " + TimeStdHour + ":" + TimeStdZeroMinute + "." + TimeStdZeroSecond

	// ==
	t.Run("normal", func(t *testing.T) {
		timeStringA := "2000/01/01 10:00.00"
		a, _ := ParseTime(timeStringA, layout)
		timeStringB := "2000/01/01 10:00.00"
		b, _ := ParseTime(timeStringB, layout)
		result := CompareTime(a, "==", b)
		assert.Equal(t, result, true)
	})
	t.Run("normal", func(t *testing.T) {
		timeStringA := "2000/01/01 10:00.00"
		a, _ := ParseTime(timeStringA, layout)
		timeStringB := "2001/01/01 10:00.00"
		b, _ := ParseTime(timeStringB, layout)
		result := CompareTime(a, "==", b)
		assert.Equal(t, result, false)
	})

	// >
	t.Run("normal", func(t *testing.T) {
		timeStringA := "2000/01/01 10:00.00"
		a, _ := ParseTime(timeStringA, layout)
		timeStringB := "2000/01/01 10:00.01"
		b, _ := ParseTime(timeStringB, layout)
		result := CompareTime(a, ">", b)
		assert.Equal(t, result, false)
	})
	t.Run("normal", func(t *testing.T) {
		timeStringA := "2000/01/01 10:00.01"
		a, _ := ParseTime(timeStringA, layout)
		timeStringB := "2000/01/01 10:00.00"
		b, _ := ParseTime(timeStringB, layout)
		result := CompareTime(a, ">", b)
		assert.Equal(t, result, true)
	})

	// <
	t.Run("normal", func(t *testing.T) {
		timeStringA := "2000/01/01 10:00.01"
		a, _ := ParseTime(timeStringA, layout)
		timeStringB := "2000/01/01 10:00.00"
		b, _ := ParseTime(timeStringB, layout)
		result := CompareTime(a, "<", b)
		assert.Equal(t, result, false)
	})
	t.Run("normal", func(t *testing.T) {
		timeStringA := "2000/01/01 10:00.00"
		a, _ := ParseTime(timeStringA, layout)
		timeStringB := "2000/01/01 10:00.01"
		b, _ := ParseTime(timeStringB, layout)
		result := CompareTime(a, "<", b)
		assert.Equal(t, result, true)
	})

	// >=
	t.Run("normal", func(t *testing.T) {
		timeStringA := "2000/01/01 10:00.00"
		a, _ := ParseTime(timeStringA, layout)
		timeStringB := "2000/01/01 10:00.00"
		b, _ := ParseTime(timeStringB, layout)
		result := CompareTime(a, ">=", b)
		assert.Equal(t, result, true)
	})
	t.Run("normal", func(t *testing.T) {
		timeStringA := "2000/01/01 10:00.01"
		a, _ := ParseTime(timeStringA, layout)
		timeStringB := "2000/01/01 10:00.00"
		b, _ := ParseTime(timeStringB, layout)
		result := CompareTime(a, ">=", b)
		assert.Equal(t, result, true)
	})
	t.Run("normal", func(t *testing.T) {
		timeStringA := "2000/01/01 10:00.00"
		a, _ := ParseTime(timeStringA, layout)
		timeStringB := "2000/01/01 10:00.01"
		b, _ := ParseTime(timeStringB, layout)
		result := CompareTime(a, ">=", b)
		assert.Equal(t, result, false)
	})

	// <=
	t.Run("normal", func(t *testing.T) {
		timeStringA := "2000/01/01 10:00.00"
		a, _ := ParseTime(timeStringA, layout)
		timeStringB := "2000/01/01 10:00.00"
		b, _ := ParseTime(timeStringB, layout)
		result := CompareTime(a, "<=", b)
		assert.Equal(t, result, true)
	})
	t.Run("normal", func(t *testing.T) {
		timeStringA := "2000/01/01 10:00.01"
		a, _ := ParseTime(timeStringA, layout)
		timeStringB := "2000/01/01 10:00.00"
		b, _ := ParseTime(timeStringB, layout)
		result := CompareTime(a, "<=", b)
		assert.Equal(t, result, false)
	})
	t.Run("normal", func(t *testing.T) {
		timeStringA := "2000/01/01 10:00.00"
		a, _ := ParseTime(timeStringA, layout)
		timeStringB := "2000/01/01 10:00.01"
		b, _ := ParseTime(timeStringB, layout)
		result := CompareTime(a, "<=", b)
		assert.Equal(t, result, true)
	})
}
