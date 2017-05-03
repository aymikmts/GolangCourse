package comma

import (
	"bytes"
)

// [ex10]bytes.Bufferを用い、再帰呼び出しを使わないバージョン
func CommaWithBuffer(s string) string {
	var buf bytes.Buffer
	// 頭から最初のコンマまでの要素数を算出
	hNum := len(s) % 3
	if hNum == 0 {
		hNum = 3
	}
	buf.WriteString(s[:hNum])

	for i, _ := range s[hNum:] {
		if i%3 == 0 {
			buf.WriteString(",")
			buf.WriteString(s[i+hNum : i+hNum+3])
		}
	}
	return buf.String()
}
