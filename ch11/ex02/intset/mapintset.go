package intset

import (
	"bytes"
	"fmt"
	"sort"
)

// MapIntSetは負ではない小さな整数のセットです。
// そのゼロ値は空セットを示しています。
type MapIntSet struct {
	set map[int]bool
}

// Hasは負ではない値xをセットが含んでいるか否かを報告します。
func (s *MapIntSet) Has(x int) bool {
	if s.set[x] {
		return true
	}
	return false
}

// Addはセットに負ではない値xを追加します。
func (s *MapIntSet) Add(x int) {
	if x < 0 {
		return
	}
	if s.set == nil {
		s.set = make(map[int]bool)
	}
	s.set[x] = true
}

// UnionWithは、sとtの和集合をsに設定します。
func (s *MapIntSet) UnionWith(t *MapIntSet) {
	if s.set == nil || t.set == nil {
		return
	}
	for k := range t.set {
		if !s.set[k] {
			s.set[k] = true
		}
	}
}

// Stringは"{1 2 3}"の形式の文字列としてセットを返します。
func (s *MapIntSet) String() string {
	if s.set == nil {
		s.set = make(map[int]bool)
	}
	var buf bytes.Buffer
	buf.WriteByte('{')

	// setの要素をソート
	var keys []int
	for k := range s.set {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, v := range keys {
		if buf.Len() > len("{") {
			buf.WriteByte(' ')
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte('}')
	return buf.String()
}
