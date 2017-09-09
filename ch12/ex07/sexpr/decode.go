package sexpr

import (
	"fmt"
	"io"
	"reflect"
	"text/scanner"
)

// Decoderは入力からS式の読み込みとデコードを行います。
type Decoder struct {
	r io.Reader
}

// NewDecoderは新しいデコーダを返します。
func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{r: r}
}

func (dec *Decoder) Decode(out interface{}) (err error) {
	lex := &lexer{scan: scanner.Scanner{Mode: scanner.GoTokens}}
	lex.scan.Init(dec.r)
	lex.next() // 最初のトークンを取得する
	defer func() {
		// 注意: これは理想的なエラー処理の例ではありません。
		if x := recover(); x != nil {
			err = fmt.Errorf("error at %s: %v", lex.scan.Position, x)
		}
	}()
	read(lex, reflect.ValueOf(out).Elem())
	return nil
}
