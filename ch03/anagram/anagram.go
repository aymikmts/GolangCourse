// Anagramは、2つの文字列が互いにアナグラムになっているかを報告します。
// "."や"?"などの記号も文字として扱います。
// 大文字・小文字は同じものとして扱います。
package anagram

import (
	"sort"
	"strings"
)

// 文字列をソートするためのインタフェース
type sortRunes []rune

func (s sortRunes) Len() int {
	return len(s)
}

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// 文字列をソートする
func SortString(s string) string {
	b := []rune(s)
	sort.Sort(sortRunes(b))
	return string(b)
}

// 2つの文字列がアナグラムになっているかどうかを返す
func IsAnagram(s1, s2 string) bool {
	// 文字を小文字にし、昇順に並び替える
	str1 := SortString(strings.ToLower(s1))
	str2 := SortString(strings.ToLower(s2))

	if str1 == str2 {
		return true
	} else {
		return false
	}
}
