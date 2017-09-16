package equal

import (
	"reflect"
	"unsafe"
)

// IsCycle は引数が循環しているデータ構造であるかどうかを報告します。
func IsCycle(x interface{}) bool {
	seen := make(map[unsafe.Pointer]bool)
	return isCycle(reflect.ValueOf(x), seen)
}

func isCycle(x reflect.Value, seen map[unsafe.Pointer]bool) bool {
	if !x.IsValid() {
		return false
	}

	// 循環の検査
	if x.CanAddr() {
		xptr := unsafe.Pointer(x.UnsafeAddr())
		if seen[xptr] {
			return true // すでに見た
		}
		seen[xptr] = true
	}

	switch x.Kind() {
	case reflect.Ptr, reflect.Interface:
		if isCycle(x.Elem(), seen) {
			return true
		}
		return false

	case reflect.Array, reflect.Slice:
		for i := 0; i < x.Len(); i++ {
			if isCycle(x.Index(i), seen) {
				return true
			}
		}
		return false

	case reflect.Struct:
		for i := 0; i < x.NumField(); i++ {
			if isCycle(x.Field(i), seen) {
				return true
			}
		}
		return false

	case reflect.Map:
		for _, key := range x.MapKeys() {
			if isCycle(x.MapIndex(key), seen) {
				return true
			}
		}
		return true
	}
	return false
}
