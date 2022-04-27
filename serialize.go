package zonotools

import (
	"encoding/gob"
	"os"
)

// Export 指定したパスにシリアライズしたオブジェクトデータを出力します。
// struct の private field は通常の Encoder では参照できず出力もされません。
// private field 含んで Encode/Decode する場合は次を参照し、独自 Encoder を実装すればよいです。
// Ref. [データのやりとりに gob を使う - Qiita](https://qiita.com/delphinus/items/67a796cb7876132a1ec0)
func Export(filePath string, p interface{}) error {
	f, err := os.Open(filePath)
	if os.IsExist(err) {
		err = os.Remove(filePath)
		if err != nil {
			return err
		}
	}
	f, err = os.Create(filePath)
	if err != nil {
		return err
	}
	defer func() {
		_ = f.Close()
	}()
	return gob.NewEncoder(f).Encode(p)
}

// Import 指定したパスからシリアライズされたオブジェクトデータを入力します。
// 第 2 引数 `p` がオブジェクトデータのマッピングされるもととなります。
func Import(filePath string, p interface{}) error {
	f, err := os.Open(filePath)
	if os.IsNotExist(err) {
		return nil
	} else if err != nil {
		return err
	}
	defer func() {
		_ = f.Close()
	}()
	return gob.NewDecoder(f).Decode(p)
}
