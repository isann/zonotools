package zonotools

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

type Hoge struct {
	Aaa string `json:"aaa"`
	Bbb string `json:"bbb"`
	ccc string `json:"ccc"`
}

func NewHoge(aaa string, bbb string, ccc string) *Hoge {
	return &Hoge{Aaa: aaa, Bbb: bbb, ccc: ccc}
}

func TestExport(t *testing.T) {
	t.Run("normal", func(t *testing.T) {
		hoge := NewHoge("0001", "0002", "0003")
		filePath := "/tmp/foobar"
		err := Export(filePath, hoge)
		if err != nil {
			assert.Error(t, err)
		}
		_, err = os.Stat(filePath)
		if err != nil {
			assert.Error(t, err)
		}
	})
}

func TestImport(t *testing.T) {
	t.Run("normal struct", func(t *testing.T) {
		hoge := NewHoge("0001", "0002", "0003")
		filePath := "/tmp/foobar"
		err := Export(filePath, hoge)
		if err != nil {
			panic(err)
		}
		_, err = os.Stat(filePath)
		if err != nil {
			panic(err)
		}
		var fuga Hoge
		err = Import(filePath, &fuga)
		if err != nil {
			assert.Error(t, err)
		}
		assert.Equal(t, hoge.Aaa, fuga.Aaa)
		assert.Equal(t, hoge.Bbb, fuga.Bbb)
		// private field のため初期値
		assert.Equal(t, "", fuga.ccc)
	})
	t.Run("normal map", func(t *testing.T) {
		type strmap map[string]string
		hoge := strmap{"001": "a", "002": "b", "003": "c"}
		filePath := "/tmp/foobar"
		err := Export(filePath, hoge)
		if err != nil {
			panic(err)
		}
		_, err = os.Stat(filePath)
		if err != nil {
			panic(err)
		}
		var fuga strmap
		err = Import(filePath, &fuga)
		if err != nil {
			assert.Error(t, err)
		}
		for k, v := range hoge {
			assert.Equal(t, v, fuga[k])
		}
	})
	t.Run("normal list", func(t *testing.T) {
		type strlist []string
		hoge := strlist{"001", "a", "002", "b", "003", "c"}
		filePath := "/tmp/foobar"
		err := Export(filePath, hoge)
		if err != nil {
			panic(err)
		}
		_, err = os.Stat(filePath)
		if err != nil {
			panic(err)
		}
		var fuga strlist
		err = Import(filePath, &fuga)
		if err != nil {
			assert.Error(t, err)
		}
		for i, v := range hoge {
			assert.Equal(t, v, fuga[i])
		}
	})
}
