package charcount

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

// CharTypeCountはUnicode分類ごとに数を数えます。
func CharTypeCount(in *bufio.Reader) (*map[string]int, int) {
	charCounts := make(map[string]int) // Unicode分類それぞれの数
	invalid := 0                       // 不正なUTF-8文字の数

	for {
		r, n, err := in.ReadRune() // rune, nbytes, errorを返す
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		// 分類ごとに数をインクリメント
		charCounts[charType(r)]++
	}

	return &charCounts, invalid
}

// charTypeはUnicode分類を返します。
func charType(r rune) string {
	// 制御文字
	if unicode.IsControl(r) {
		return "control"
	}

	// 文字
	if unicode.IsLetter(r) {
		return "letter"
	}

	// 数値
	if unicode.IsNumber(r) {
		return "number"
	}

	// スペース
	if unicode.IsSpace(r) {
		return "space"
	}

	// 記号
	if unicode.IsSymbol(r) {
		return "symbol"
	}

	return "other"
}
