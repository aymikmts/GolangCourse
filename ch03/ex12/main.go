// Ex12は2つの文字列が互いにアナグラムになっているかを報告します。
// 実装は"anagram/anagram.go"
package main

import (
	"fmt"
	"os"

	"../anagram"
	//"GolangCourse/ch03/anagram"
)

func main() {
	fmt.Printf("str1:%s str2:%s IsAnagram:%v\n", os.Args[1], os.Args[2], anagram.IsAnagram(os.Args[1], os.Args[2]))
}
