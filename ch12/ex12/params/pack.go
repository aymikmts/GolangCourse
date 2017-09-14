package params

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// DefaultMaxResults はMaxResultsのデフォルト値です。
const DefaultMaxResults = 10

// Pack はptrが指す構造体のフィールドからURLを生成し返します。
func Pack(ptr interface{}) (string, error) {

	// 実効的な名前をキーとするフィールドのマップを構築する。
	fields := make(map[string]reflect.Value)
	v := reflect.ValueOf(ptr)
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i) // reflect.StructField
		tag := fieldInfo.Tag           // reflect.StructTag
		name := tag.Get("http")
		if name == "" {
			name = strings.ToLower(fieldInfo.Name)
		}
		fields[name] = v.Field(i)
	}

	// フィールドごとにパラメータを抽出し、文字列を作成する
	elms := make(map[string]string) // フィールドごとの文字列
	for name, value := range fields {
		// パラメータ値を読み込み、[]stringにまとめる
		var slice []string
		if err := extract(&slice, value); err != nil {
			return "", err
		}

		// sliceを&で連結
		var buf bytes.Buffer
		for i, s := range slice {
			// maxパラメータはデフォルト値を考慮する
			if name == "max" {
				val, err := strconv.Atoi(s)
				if err != nil {
					return "", err
				}
				if val == DefaultMaxResults {
					continue
				}
			}
			if i > 0 {
				buf.WriteRune('&')
			}
			fmt.Fprintf(&buf, "%s=%s", name, s)
		}
		elms[name] = string(buf.Bytes())
	}

	// 各パラメータの文字列を&で連結
	var buf bytes.Buffer
	n := 0
	for _, elm := range elms {
		if len(elm) == 0 {
			continue
		}
		if n > 0 && n < len(elms)-1 {
			buf.WriteRune('&')
		}
		buf.WriteString(elm)
		n++
	}

	return string(buf.Bytes()), nil
}

func extract(ps *[]string, v reflect.Value) error {
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		*ps = append(*ps, strconv.Itoa(int(v.Int())))

	case reflect.String:
		*ps = append(*ps, v.String())

	case reflect.Array, reflect.Slice:
		for i := 0; i < v.Len(); i++ {
			if err := extract(ps, v.Index(i)); err != nil {
				return err
			}
		}

	case reflect.Bool:
		if v.Bool() {
			*ps = append(*ps, strconv.FormatBool(v.Bool()))
		}

	default:
		return fmt.Errorf("type \"%s\" is not implemented yet", v.Type())
	}
	return nil
}
