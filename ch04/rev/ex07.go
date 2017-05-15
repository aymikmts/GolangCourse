package rev

// ReverseUnicodeByteはUTF-8でエンコードされた文字列を表す[]byteスライスを直接逆順に並び替えます。
func ReverseUnicodeByte(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+3, j-3 {
		s[i], s[i+1], s[i+2], s[j-2], s[j-1], s[j] = s[j-2], s[j-1], s[j], s[i], s[i+1], s[i+2]
	}
}
