package comma

import (
	"bytes"
	"fmt"
	"strings"
)

// [ex11]符号記号を持つ浮動小数点数にカンマを挿入します。
func CommaSignedFloat(s string) string {
	var buf bytes.Buffer
	var sIdx int              // 符号を取り除いた文字列の最初のインデックス番号
	var numIdx int            // 0以外の数値文字列の頭のインデックス番号
	var dotIdx int            // 小数点のインデックス番号
	var eNum int              // 累乗数
	var floatBuf bytes.Buffer // 浮動小数点表現時の、小数点以下の部分文字列

	// 符号があればbufに追記し、整数部分の頭のインデックス番号を記録
	if s[0] == '-' || s[0] == '+' {
		buf.WriteString(string(s[0]))
		sIdx = 1
	}

	// 文字列から"."を探す
	dotIdx = strings.LastIndex(s, ".")

	if dotIdx >= 0 {
		// "."があったとき(小数点を含むとき)
		numIdx, eNum, floatBuf = calcNumsFloat(sIdx, dotIdx, s)
	} else {
		// "."がなかったとき(小数点を含まないとき)
		dotIdx = len(s)
		numIdx, eNum, floatBuf = calcNumsInt(sIdx, s)
	}

	// 表示する文字列の連結
	buf.WriteString(string(s[numIdx]))     // 整数部
	buf.WriteString(".")                   // 小数点
	buf.WriteString(floatBuf.String())     // 小数点以下
	eNumStr := fmt.Sprintf("e%+03d", eNum) // 指数部
	buf.WriteString(eNumStr)

	return buf.String()
}

func calcNumsFloat(sIdx int, dotIdx int, s string) (numIdx, eNum int, floatBuf bytes.Buffer) {
	for i, v := range s[sIdx:dotIdx] {
		// "."より左に"0"以外があった場合
		if v != '0' {
			numIdx = sIdx
			eNum = dotIdx - numIdx - 1
			floatBuf.WriteString(s[numIdx+1 : dotIdx])
			floatBuf.WriteString(s[dotIdx+1:])
			return
		}

		// "."より左が全て"0"の場合
		if i == dotIdx-sIdx-1 {
			for j, w := range s[dotIdx+1:] {
				if w != '0' {
					numIdx = dotIdx + 1 + j
					eNum = dotIdx - numIdx
					floatBuf.WriteString(s[numIdx+1:])
					return
				}
			}
		}
	}
	return
}

func calcNumsInt(sIdx int, s string) (numIdx, eNum int, floatBuf bytes.Buffer) {
	dotIdx := len(s)
	numIdx = sIdx

	for i, v := range s[sIdx:] {
		if v != '0' {
			numIdx = sIdx + i
			break
		}
	}

	eNum = dotIdx - numIdx - 1
	floatBuf.WriteString(s[numIdx+1:])
	return
}
