package sexpr

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
	"strconv"
	"text/scanner"
)

// Decoderは入力からS式の読み込みとデコードを行います。
type Decoder struct {
	r   io.Reader
	lex *lexer
}

// Tokenは次のtoken typesをもったインタフェースです。
// Symbol, String, Int, StartList, EndList
type Token interface{}

type Symbol struct {
	Name string
}

type String struct {
	Value string
}

type Int struct {
	Value int
}

type StartList struct{}

type EndList struct{}

// NewDecoderは新しいデコーダを返します。
func NewDecoder(r io.Reader) *Decoder {
	var dec Decoder
	dec.r = r
	dec.lex = &lexer{scan: scanner.Scanner{Mode: scanner.GoTokens}}
	dec.lex.scan.Init(dec.r)
	dec.lex.next() // 最初のトークンを取得する
	return &dec
}

func (dec *Decoder) Decode(out interface{}) (err error) {
	defer func() {
		// 注意: これは理想的なエラー処理の例ではありません。
		if x := recover(); x != nil {
			err = fmt.Errorf("error at %s: %v", dec.lex.scan.Position, x)
		}
	}()
	read(dec.lex, reflect.ValueOf(out).Elem())
	return nil
}

// ex09で追加
func (dec *Decoder) Token() Token {
	switch dec.lex.token {
	case scanner.Ident:
		name := dec.lex.text()
		dec.lex.next()
		return Symbol{Name: name}

	case scanner.String:
		s, _ := strconv.Unquote(dec.lex.text())
		dec.lex.next()
		return String{Value: s}

	case scanner.Int:
		i, _ := strconv.Atoi(dec.lex.text())
		dec.lex.next()
		return Int{Value: i}

	case '(':
		dec.lex.next()
		return StartList{}

	case ')':
		dec.lex.next()
		return EndList{}
	}
	panic(fmt.Sprintf("unexpected token %v", dec.lex.token))
}

// UnmarshalはS式のデータをパースしてnilではないポインタ
// outにアドレスが入っている変数に移しかえます。
func Unmarshal(data []byte, out interface{}) (err error) {
	dec := NewDecoder(bytes.NewReader(data))
	return dec.Decode(out)
}

// lexerは、Scanから返された最後のトークンを記録しておく、スキャナを包んだヘルパー型です。
type lexer struct {
	scan  scanner.Scanner
	token rune
}

func (lex *lexer) next()        { lex.token = lex.scan.Scan() }
func (lex *lexer) text() string { return lex.scan.TokenText() }

func (lex *lexer) consume(want rune) {
	if lex.token != want {
		panic(fmt.Sprintf("got %q, want %q", lex.text(), want))
	}
	lex.next()
}

func read(lex *lexer, v reflect.Value) {
	switch lex.token {
	case scanner.Ident:
		// 唯一の正当な識別子は"nil"と
		// 構造体のフィールド名です。
		switch lex.text() {
		case "nil":
			v.Set(reflect.Zero(v.Type()))
			lex.next()
			return

		// ex03で追加
		case "t":
			v.SetBool(true)
			lex.next()
			return

		}
	case scanner.String:
		s, _ := strconv.Unquote(lex.text()) // 注意: エラーを無視している
		v.SetString(s)
		lex.next()
		return
	case scanner.Int:
		i, _ := strconv.Atoi(lex.text()) // 注意: エラーを無視している
		v.SetInt(int64(i))
		lex.next()
		return

		// ex03で追加
	case scanner.Float:
		switch v.Kind() {
		case reflect.Float32:
			f, _ := strconv.ParseFloat(lex.text(), 32)
			v.SetFloat(f)
		case reflect.Float64:
			f, _ := strconv.ParseFloat(lex.text(), 64)
			v.SetFloat(f)
		default:
			panic(fmt.Sprintf("unexpected type %v", v.Kind()))
		}
		lex.next()
		return

	case '#': // #C(float, float)
		lex.next()
		lex.next()
		lex.next()
		real := lex.text()
		lex.next()
		imag := lex.text()
		lex.next()
		lex.consume(')')

		var bitSize int
		switch v.Kind() {
		case reflect.Complex64:
			bitSize = 32
		case reflect.Complex128:
			bitSize = 64
		default:
			panic(fmt.Sprintf("unexpected type: %v", v.Kind()))
		}

		cReal, _ := strconv.ParseFloat(real, bitSize)
		cImag, _ := strconv.ParseFloat(imag, bitSize)
		v.SetComplex(complex(cReal, cImag))
		return

	case '(':
		lex.next()
		readList(lex, v)
		lex.next() // ')'を消費する
		return
	}
	panic(fmt.Sprintf("unexpected token %q", lex.text()))
}

func readList(lex *lexer, v reflect.Value) {
	switch v.Kind() {
	case reflect.Array: // (item ...)
		for i := 0; !endList(lex); i++ {
			read(lex, v.Index(i))
		}
	case reflect.Slice: // (item ...)
		for !endList(lex) {
			item := reflect.New(v.Type().Elem()).Elem()
			read(lex, item)
			v.Set(reflect.Append(v, item))
		}
	case reflect.Struct: // ((name value) ...)
		for !endList(lex) {
			lex.consume('(')
			if lex.token != scanner.Ident {
				panic(fmt.Sprintf("got token %q, want field name",
					lex.text()))
			}
			name := lex.text()
			lex.next()
			read(lex, v.FieldByName(name))
			lex.consume(')')
		}
	case reflect.Map: // ((key value) ...)
		v.Set(reflect.MakeMap(v.Type()))
		for !endList(lex) {
			lex.consume('(')
			key := reflect.New(v.Type().Key()).Elem()
			read(lex, key)
			value := reflect.New(v.Type().Elem()).Elem()
			read(lex, value)
			v.SetMapIndex(key, value)
			lex.consume(')')
		}
	default:
		panic(fmt.Sprintf("cannot decode list into %v", v.Type()))
	}
}

func endList(lex *lexer) bool {
	switch lex.token {
	case scanner.EOF:
		panic("end of file")
	case ')':
		return true
	}
	return false
}
