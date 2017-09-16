package sexpr

import (
	"bytes"
	"fmt"
	"reflect"
)

// MarshalはGoの値をS式形式でエンコードします。
func Marshal(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	if err := encode(&buf, reflect.ValueOf(v)); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func encode(buf *bytes.Buffer, v reflect.Value) error {
	switch v.Kind() {
	case reflect.Invalid:
		buf.WriteString("nil")

	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		fmt.Fprintf(buf, "%d", v.Int())

	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		fmt.Fprintf(buf, "%d", v.Uint())

	case reflect.String:
		fmt.Fprintf(buf, "%q", v.String())

	case reflect.Ptr:
		return encode(buf, v.Elem())

	case reflect.Array, reflect.Slice: // (value ...)
		buf.WriteByte('(')
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				buf.WriteByte(' ')
			}
			if err := encode(buf, v.Index(i)); err != nil {
				return err
			}
		}
		buf.WriteByte(')')

	case reflect.Struct: // ((name value) ...)

		if !isZeroValue(v) {

			first := true
			buf.WriteByte('(')
			for i := 0; i < v.NumField(); i++ {

				if isZeroValue(v.Field(i)) {
					continue
				}

				if i > 0 && !first {
					buf.WriteByte(' ')
				} else {
					first = false
				}

				// ex13で追加
				fieldInfo := v.Type().Field(i)
				tag := fieldInfo.Tag
				name := tag.Get("sexpr")
				if name == "" {
					name = fieldInfo.Name
				}

				fmt.Fprintf(buf, "(%s ", name)
				if err := encode(buf, v.Field(i)); err != nil {
					return err
				}
				buf.WriteByte(')')
			}
			buf.WriteByte(')')
		}

	case reflect.Map: // ((key value) ...)
		buf.WriteByte('(')
		for i, key := range v.MapKeys() {
			if i > 0 {
				buf.WriteByte(' ')
			}
			buf.WriteByte('(')
			if err := encode(buf, key); err != nil {
				return err
			}
			buf.WriteByte(' ')
			if err := encode(buf, v.MapIndex(key)); err != nil {
				return err
			}
			buf.WriteByte(')')
		}
		buf.WriteByte(')')

	// ex03で追加
	case reflect.Bool:
		if v.Bool() == true {
			buf.WriteString("t")
		} else {
			buf.WriteString("nil")
		}
		return nil
	case reflect.Float32, reflect.Float64:
		fmt.Fprintf(buf, "%f", v.Float())
		return nil
	case reflect.Complex64, reflect.Complex128:
		fmt.Fprintf(buf, "#C(%f %f)", real(v.Complex()), imag(v.Complex()))
		return nil
	case reflect.Interface:
		if v.IsNil() {
			buf.WriteString("nil")
		} else {
			buf.WriteByte('(')
			fmt.Fprintf(buf, "\"%s\" ", v.Elem().Type())
			encode(buf, v.Elem())
			buf.WriteByte(')')
		}
		return nil

	default: // chan, func
		return fmt.Errorf("unsupported type: %s", v.Type())
	}
	return nil
}

func isZeroValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		return v.Int() == 0

	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0

	case reflect.String, reflect.Array, reflect.Slice:
		return v.Len() == 0

	case reflect.Ptr, reflect.Interface:
		return v.IsNil()

	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			if !isZeroValue(v.Field(i)) {
				return false
			}
		}
		return true

	case reflect.Map:
		for _, key := range v.MapKeys() {
			if !isZeroValue(v.MapIndex(key)) {
				return false
			}
		}
		return true

	case reflect.Bool:
		return !v.Bool()

	case reflect.Float32, reflect.Float64:
		return v.Float() == 0

	case reflect.Complex64, reflect.Complex128:
		return v.Complex() == 0
	}
	return false
}
