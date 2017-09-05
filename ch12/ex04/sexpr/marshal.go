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

// ex04で追加
// MarshalWithIndentはGoの値をインデント付きS式形式でエンコードします。
func MarshalWithIndent(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	if err := encodeWithIndent(&buf, reflect.ValueOf(v)); err != nil {
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
		buf.WriteByte('(')
		for i := 0; i < v.NumField(); i++ {
			if i > 0 {
				buf.WriteByte(' ')
			}
			fmt.Fprintf(buf, "(%s ", v.Type().Field(i).Name)
			if err := encode(buf, v.Field(i)); err != nil {
				return err
			}
			buf.WriteByte(')')
		}
		buf.WriteByte(')')

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

func encodeWithIndent(buf *bytes.Buffer, v reflect.Value) error {
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
		return encodeWithIndent(buf, v.Elem())

	case reflect.Array, reflect.Slice: // (value ...)
		buf.WriteByte('(')
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				if i != v.Len() {
					buf.WriteString("\n\t")
				} else {
					buf.WriteByte(' ')
				}
			}
			if err := encodeWithIndent(buf, v.Index(i)); err != nil {
				return err
			}
		}
		buf.WriteByte(')')

	case reflect.Struct: // ((name value) ...)
		buf.WriteByte('(')
		for i := 0; i < v.NumField(); i++ {
			if i > 0 {
				buf.WriteByte(' ')
			}
			fmt.Fprintf(buf, "(%s ", v.Type().Field(i).Name)
			if err := encodeWithIndent(buf, v.Field(i)); err != nil {
				return err
			}
			if i != v.NumField()-1 {
				buf.WriteString(")\n")
			} else {
				buf.WriteByte(')')
			}
		}
		buf.WriteByte(')')

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
			if err := encodeWithIndent(buf, v.MapIndex(key)); err != nil {
				return err
			}

			if i != v.Len()-1 {
				buf.WriteString(")\n\t")
			} else {
				buf.WriteByte(')')
			}
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
			encodeWithIndent(buf, v.Elem())
			buf.WriteByte(')')
		}
		return nil

	default: // chan, func
		return fmt.Errorf("unsupported type: %s", v.Type())
	}
	return nil
}
