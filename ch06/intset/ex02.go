package intset

// AddAllはセットに負ではない値xを追加します。
func (s *IntSet) AddAll(vals ...int) {
	for _, val := range vals {
		word, bit := val/64, uint(val%64)
		for word >= len(s.words) {
			s.words = append(s.words, 0)
		}
		s.words[word] |= 1 << bit
	}
}
