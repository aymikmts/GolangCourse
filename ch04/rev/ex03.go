package rev

// ReverseByPointerは大きさ[6]のint配列を直接逆順に並び替えます。
func ReverseByPointer(s *[6]int) {
	for i, j := 0, len(*s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
