package main

import (
	"bytes"
	"fmt"
)

const UINTSIZE = 32 << (^uint(0) >> 63)

// IntSetは負ではない小さな整数のセットです。
// そのゼロ値は空セットを示しています。
type IntSet struct {
	words []uint
}

// Hasは負ではない値xをセットが含んでいるか否かを報告します。
func (s *IntSet) Has(x int) bool {
	word, bit := x/UINTSIZE, uint(x%UINTSIZE)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Addはセットに負ではない値xを追加します。
func (s *IntSet) Add(x int) {
	word, bit := x/UINTSIZE, uint(x%UINTSIZE)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWithは、sとtの和集合をsに設定します。
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// Stringは"{1 2 3}"の形式の文字列としてセットを返します。
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < UINTSIZE; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", UINTSIZE*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

// Lenは要素数を返します。
func (s *IntSet) Len() int {
	var ret int
	for _, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < UINTSIZE; j++ {
			if word&(1<<uint(j)) != 0 {
				ret++
			}
		}
	}
	return ret
}

// Removeはセットからxを取り除きます。
func (s *IntSet) Remove(x int) {
	word, bit := x/UINTSIZE, uint(x%UINTSIZE)
	s.words[word] &^= 1 << bit
}

// Clearはセットからすべての要素を取り除きます。
func (s *IntSet) Clear() {
	s.words = []uint{}
}

// Copyはセットのコピーを返します。
func (s *IntSet) Copy() *IntSet {
	var ret IntSet
	for _, word := range s.words {
		ret.words = append(ret.words, word)
	}
	return &ret
}

// AddAllはセットに負ではない値xを追加します。
func (s *IntSet) AddAll(vals ...int) {
	for _, val := range vals {
		word, bit := val/UINTSIZE, uint(val%UINTSIZE)
		for word >= len(s.words) {
			s.words = append(s.words, 0)
		}
		s.words[word] |= 1 << bit
	}
}

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

func (s *IntSet) Elems() []int {
	var elems []int

	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < UINTSIZE; j++ {
			if word&(1<<uint(j)) != 0 {
				elems = append(elems, UINTSIZE*i+j)
			}
		}
	}

	return elems
}

func main() {
	fmt.Println(UINTSIZE)
}
