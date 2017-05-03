// Commaは負ではない10進表記整数文字列にカンマを挿入します。
package comma

func Comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return Comma(s[:n-3]) + "," + s[n-3:]
}
