package intset

// IntersectWithは、sとtの積集合をsに設定します。
func (s *IntSet) IntersectWith(t *IntSet) {
	for i := 0; i < len(s.words); i++ {
		if i < len(t.words) {
			s.words[i] &= t.words[i]
		} else {
			s.words[i] = 0
		}
	}
}

// DifferenceWithは、sとtの差集合をsに設定します。
func (s *IntSet) DifferenceWith(t *IntSet) {
	for i := 0; i < len(s.words); i++ {
		if i < len(t.words) {
			s.words[i] &^= t.words[i]
		}
	}
}

// SymmetricDifferenceは、sとtの対称差集合をsに設定します。
func (s *IntSet) SymmetricDifference(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] ^= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}
