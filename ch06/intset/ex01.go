package intset

// Lenは要素数を返します。
func (s *IntSet) Len() int {
	var ret int
	for _, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				ret++
			}
		}
	}
	return ret
}

// Removeはセットからxを取り除きます。
func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	s.words[word] &^= 1 << bit
}

// Clearはセットからすべての要素を取り除きます。
func (s *IntSet) Clear() {
	s.words = []uint64{}
}

// Copyはセットのコピーを返します。
func (s *IntSet) Copy() *IntSet {
	var ret IntSet
	for _, word := range s.words {
		ret.words = append(ret.words, word)
	}
	return &ret
}
