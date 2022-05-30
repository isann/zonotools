package zonotools

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestOrgStruct struct {
	Name string
	Num  int
	Foo  interface{}
}

type TestOrgStruct2 struct {
	AName string
	ANum  int
	AFoo  interface{}
}

type TestCopyStruct struct {
	Name string
	Num  int
	Bar  interface{}
}

func TestMapper(t *testing.T) {
	//type args struct {
	//	org  interface{}
	//	copy interface{}
	//}
	//tests := []struct {
	//	name string
	//	args args
	//}{
	//	{name: "normal", args: args{&TestOrgStruct{"a", 1, map[string]string{}}, &TestCopyStruct{}}},
	//	{name: "fail", args: args{TestOrgStruct{"a", 1, map[string]string{}}, TestCopyStruct{}}},
	//}
	//for _, tt := range tests {
	//	t.Run(tt.name, func(t *testing.T) {
	//		Mapper(tt.args.org, tt.args.copy)
	//		assert.Equal(t, tt.args.org.(TestOrgStruct).Name, tt.args.copy.(TestCopyStruct).Name)
	//		assert.Equal(t, tt.args.org.(TestOrgStruct).Num, tt.args.copy.(TestCopyStruct).Num)
	//		assert.NotEqual(t, tt.args.org.(TestOrgStruct).Foo, tt.args.copy.(TestCopyStruct).Bar)
	//	})
	//}
	t.Run("normal", func(t *testing.T) {
		a1 := &TestOrgStruct{"a", 1, map[string]string{}}
		a2 := &TestCopyStruct{}
		Mapper(a1, a2)
		assert.Equal(t, a1.Name, a2.Name)
		assert.Equal(t, a1.Num, a2.Num)
		assert.NotEqual(t, a1.Foo, a2.Bar)
	})
	// Generics により参照型の受け渡し以外ビルドエラーとなるので次のテストコードは不要になった
	//t.Run("fail-01", func(t *testing.T) {
	//	a1 := TestOrgStruct{"a", 1, map[string]string{}}
	//	a2 := TestCopyStruct{}
	//	Mapper(a1, a2)
	//	assert.NotEqual(t, a1.Name, a2.Name)
	//	assert.NotEqual(t, a1.Num, a2.Num)
	//	assert.NotEqual(t, a1.Foo, a2.Bar)
	//})
	//t.Run("fail-02", func(t *testing.T) {
	//	a1 := TestOrgStruct2{"a", 1, map[string]string{}}
	//	a2 := TestCopyStruct{}
	//	Mapper(a1, a2)
	//	assert.NotEqual(t, a1.AName, a2.Name)
	//	assert.NotEqual(t, a1.ANum, a2.Num)
	//	assert.NotEqual(t, a1.AFoo, a2.Bar)
	//})
}

func TestIsExistsKey(t *testing.T) {
	t.Run("normal", func(t *testing.T) {
		m := map[string]string{"abc": "123"}
		ret := IsExistsKey(m, "abc")
		assert.Equal(t, true, ret)
	})
	t.Run("normal", func(t *testing.T) {
		m := map[string]string{"abc": "123"}
		ret := IsExistsKey(m, "aaa")
		assert.False(t, ret)
	})
}

func TestIsExistKey(t *testing.T) {
	t.Run("normal", func(t *testing.T) {
		m := map[string]string{"abc": "123"}
		val, ret := IsExistKey(m, "abc")
		assert.True(t, ret)
		assert.Equal(t, "123", val)
	})
	t.Run("normal", func(t *testing.T) {
		m := map[string]int64{"abc": int64(123)}
		val, ret := IsExistKey(m, "abc")
		assert.True(t, ret)
		assert.Equal(t, int64(123), val)
	})
}
