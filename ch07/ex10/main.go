// Ex10は列sが回分かどうかを報告します。
package main

import (
	"fmt"
	"log"
	"os"
	"sort"
)

type runes []rune

func (r runes) Len() int           { return len(r) }
func (r runes) Less(i, j int) bool { return r[i] < r[j] }
func (r runes) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }

func IsPalindrome(s sort.Interface) bool {
	length := s.Len()
	for i := 0; i < length/2; i++ {
		if !s.Less(i, length-i-1) && !s.Less(length-i-1, i) {
			continue
		}
		return false
	}
	return true
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("no args. input a text.")
	}
	str := []rune(os.Args[1])
	fmt.Printf("input: %q\n", str)
	fmt.Printf("IsPalindrome: %v\n", IsPalindrome(runes(str)))
}
