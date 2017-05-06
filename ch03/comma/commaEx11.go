package comma

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// [ex11]符号記号を持つ浮動小数点数にカンマを挿入します。
func CommaSignedFloat(s string) string {
	var buf bytes.Buffer
	var sIdx, dotIdx, eIdx int
	var floatStr string

	// 浮動小数点表記を指数なし表記に変換
	num, err := strconv.ParseFloat(s, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error:%v input:%v\n", err, s)
		return s
	}
	numStr := fmt.Sprintf("%f", num)

	// 符号があればbufに追記し、
	// 整数部分のインデックス番号を記録
	if s[0] == '-' || s[0] == '+' {
		buf.WriteString(string(numStr[0]))
		sIdx = 1
	}

	// 文字列から"."を探し、
	// そのインデックス番号を記録
	dotIdx = strings.LastIndex(numStr, ".")

	// 文字列の一番後ろの数値のインデックス番号を記録
	eIdx = dotIdx
	for i := len(numStr) - 1; i > 0; i-- {
		if numStr[i] == '0' {
			continue
		} else {
			eIdx = i + 1
			break
		}
	}

	// 小数点以下の文字列を抽出
	floatStr = numStr[dotIdx:eIdx]

	// 整数部分にカンマを挿入
	intStr := CommaWithBuffer(numStr[sIdx:dotIdx])

	// 整数部分・小数点以下部分を連結
	buf.WriteString(intStr)
	if len(floatStr) != 1 {
		// 小数点以下がなかったら連結しない
		buf.WriteString(floatStr)
	}

	return buf.String()
}
