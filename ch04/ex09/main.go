// Ex09はテキストファイル内のそれぞれの単語の出現頻度を報告します。
// 実装は"charcount/ex09.go"
package main

import (
	"bufio"
	"fmt"
	"os"

	"../charcount"
)

func main() {
	input := bufio.NewScanner(os.Stdin)

	wordFreq := charcount.WordFreq(input)

	fmt.Printf("WORD\tFREQUENCY\n")
	for word, freq := range *wordFreq {
		fmt.Printf("%s\t%3.2f%%\n", word, freq)
	}
}
