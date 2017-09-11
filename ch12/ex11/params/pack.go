package params

import "fmt"

// Pack はptrが指す構造体のフィールドからURLを生成し返します。
func Pack(ptr interface{}) (string, error) {
	return fmt.Sprintf("%v", ptr), nil
}
