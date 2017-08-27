package display

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

var Depth int = 3

func Display(name string, x interface{}) {
	fmt.Printf("Display %s (%T):\n", name, x)
	count := 0
	display(name, reflect.ValueOf(x), &count)
}

func display(path string, v reflect.Value, count *int) {
	if *count >= Depth {
		return
	}

	switch v.Kind() {
	case reflect.Invalid:
		fmt.Printf("%s = invalid\n", path)
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			display(fmt.Sprintf("%s[%d]", path, i), v.Index(i), count)
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			fieldPath := fmt.Sprintf("%s.%s", path, v.Type().Field(i).Name)
			display(fieldPath, v.Field(i), count)
		}
	case reflect.Map:
		for _, key := range v.MapKeys() {
			display(fmt.Sprintf("%s[%s]", path,
				formatAtom(key)), v.MapIndex(key), count)
		}
	case reflect.Ptr:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			display(fmt.Sprintf("(*%s)", path), v.Elem(), count)
		}
	case reflect.Interface:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			fmt.Printf("%s.type = %s\n", path, v.Elem().Type())
			display(path+".value", v.Elem(), count)
		}
	default: // basic types, channels, funcs
		fmt.Printf("%s = %s\n", path, formatAtom(v))
	}

	*count++
}

// formatAtomは値の内部構造を調べることなくその値をフォーマットします。
func formatAtom(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10)
	// ... 簡素にするために浮動小数点数と複素数は省略...
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.String:
		return strconv.Quote(v.String())
	case reflect.Chan, reflect.Func, reflect.Ptr, reflect.Slice, reflect.Map:
		return v.Type().String() + "0x" + strconv.FormatUint(uint64(v.Pointer()), 16)

	// ex01で追加
	// reflect.Struct
	case reflect.Struct:
		var elms []string
		for i := 0; i < v.NumField(); i++ {
			elm := fmt.Sprintf("%s: %v", v.Type().Field(i).Name, formatAtom(v.Field(i)))
			elms = append(elms, elm)
		}
		return v.Type().Name() + "{" + strings.Join(elms, ", ") + "}"
	// reflect.Array
	case reflect.Array:
		var elms []string
		for i := 0; i < v.Len(); i++ {
			elm := fmt.Sprintf("%v", formatAtom(v.Index(i)))
			elms = append(elms, elm)
		}
		return v.Type().String() + "{" + strings.Join(elms, ", ") + "}"

	default: // reflect.Interface
		return v.Type().String() + "value"
	}
}
