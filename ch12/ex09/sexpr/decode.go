package sexpr

import (
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
	// lex := &lexer{scan: scanner.Scanner{Mode: scanner.GoTokens}}
	// lex.scan.Init(dec.r)
	// lex.next() // 最初のトークンを取得する
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
