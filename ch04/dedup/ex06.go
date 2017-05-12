package dedup

import (
	"bytes"
	"unicode"
)

const (
	Space rune = '\u0020' // ASCIIスペース
)

// DedupSpaceは[]byteスライス内で隣接しているUnicodeスペースをASCIIスペースへ圧縮します。
func DedupSpace(s []byte) []byte {
	var ret []rune

	runeS := bytes.Runes(s) // []byteを[]runeに変換
	for i := 0; i < len(runeS); i++ {
		if unicode.IsSpace(runeS[i]) {
			// 空白文字であったら、次以降の文字が空白かどうかを見る
			j := 0
			for i+j < len(runeS) {
				// 空白でない文字以外になるまでインデックスをインクリメント
				if !unicode.IsSpace(runeS[i+j]) {
					break
				}
				j++
			}

			if j <= 1 {
				// 空白文字が連続していなかったら、元の空白文字を記録
				ret = append(ret, runeS[i])
			} else {
				// 空白文字が連続していたら、ASCIIスペース1文字を記録
				ret = append(ret, Space)
			}

			// 空白文字の次の文字を記録
			if i+j < len(runeS) {
				ret = append(ret, runeS[i+j])
			}

			i += j

		} else {
			ret = append(ret, runeS[i])
		}
	}

	return []byte(string(ret))
}
