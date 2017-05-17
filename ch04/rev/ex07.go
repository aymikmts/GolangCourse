package rev

// ReverseUnicodeByteはUTF-8でエンコードされた文字列を表す[]byteスライスを直接逆順に並び替えます。
func ReverseUnicodeByte(s []byte) {

	// 文字単位でbyte列を逆順にする
	for i := 0; i < len(s); i++ {
		// ASCII文字のとき
		if ((s[i] >> 7) ^ 0x00) == 0x00 {
			continue
		}

		// 2byte文字のとき
		if ((s[i] >> 5) ^ 0x06) == 0x00 {
			s[i], s[i+1] = s[i+1], s[i]
			i = i + 1
			continue
		}

		// 3byte文字のとき
		if ((s[i] >> 4) ^ 0x0E) == 0x00 {
			s[i], s[i+2] = s[i+2], s[i]
			i = i + 2
			continue
		}

		// 4byte文字のとき
		if ((s[i] >> 3) ^ 0x1E) == 0x00 {
			s[i], s[i+1], s[i+2], s[i+3] = s[i+3], s[i+2], s[i+1], s[i]
			i = i + 3
			continue
		}
	}

	// 全体を逆順にする
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
