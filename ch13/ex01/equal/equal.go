package equal

import (
	"math"
	"reflect"
	"unsafe"
)

// Equal はxとyが深く等しいかどうかを報告します。
func Equal(x, y interface{}) bool {
	seen := make(map[comparison]bool)
	return equal(reflect.ValueOf(x), reflect.ValueOf(y), seen)
}

type comparison struct {
	x, y unsafe.Pointer
	t    reflect.Type
}

func equal(x, y reflect.Value, seen map[comparison]bool) bool {
	if !x.IsValid() || !y.IsValid() {
		return x.IsValid() == y.IsValid()
	}
	if x.Type() != y.Type() {
		return false
	}

	// 循環の検査
	if x.CanAddr() && y.CanAddr() {
		xptr := unsafe.Pointer(x.UnsafeAddr())
		yptr := unsafe.Pointer(y.UnsafeAddr())
		if xptr == yptr {
			return true // 同一参照
		}
		c := comparison{xptr, yptr, x.Type()}
		if seen[c] {
			return true // すでに見た
		}
		seen[c] = true
	}

	switch x.Kind() {
	case reflect.Bool:
		return x.Bool() == y.Bool()

	case reflect.String:
		return x.String() == y.String()

	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		return x.Int() == y.Int()

	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return x.Uint() == y.Uint()

	case reflect.Float32, reflect.Float64:
		return equalFloat(x.Float(), y.Float())

	case reflect.Complex64, reflect.Complex128:
		return equalComplex(x.Complex(), y.Complex())

	case reflect.Chan, reflect.UnsafePointer, reflect.Func:
		return x.Pointer() == y.Pointer()

	case reflect.Ptr, reflect.Interface:
		return equal(x.Elem(), y.Elem(), seen)

	case reflect.Array, reflect.Slice:
		if x.Len() != y.Len() {
			return false
		}
		for i := 0; i < x.Len(); i++ {
			if !equal(x.Index(i), y.Index(i), seen) {
				return false
			}
		}
		return true

	case reflect.Struct:
		if x.NumField() != y.NumField() {
			return false
		}
		for i := 0; i < x.NumField(); i++ {
			if !equal(x.Field(i), y.Field(i), seen) {
				return false
			}
		}
		return true

	case reflect.Map:
		if x.Len() != y.Len() {
			return false
		}
		for _, key := range x.MapKeys() {
			if !equal(x.MapIndex(key), y.MapIndex(key), seen) {
				return false
			}
		}
		return true
	}
	panic("unreachable")
}

const Diff = 1.0e-9

func equalFloat(x, y float64) bool {
	return math.Abs(x-y) < Diff
}

func equalComplex(x, y complex128) bool {
	return equalFloat(real(x), real(y)) && equalFloat(imag(x), imag(y))

}
