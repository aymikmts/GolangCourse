// Ex08はUnicode分類ごとに数を数えます。
// 標準入力から数える対象を入力します。
// 実装は"charcount/ex08.go"
package main

import (
	"bufio"
	"fmt"
	"os"

	"../charcount"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	charCounts, invalid := charcount.CharTypeCount(in)

	fmt.Printf("charType\tcount\n")
	for i, n := range *charCounts {
		fmt.Printf("%s\t\t%d\n", i, n)
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
