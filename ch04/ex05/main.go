// Ex05は[]stringスライス内で隣接している重複を除去します。
// スライスは標準入力より入力された数値から生成します。
// 実装は"dedup/ex05.go"
package main

import (
	//"GolangCourse/ch04/dedup

	"bufio"
	"fmt"
	"os"
	"strings"

	"../dedup"
)

func main() {

	// 配列の作成
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		slice := strings.Fields(input.Text())

		fmt.Printf("input : %v\n", slice)

		// Ex05の実行
		slice = dedup.Dedup(slice)

		fmt.Printf("output: %v\n\n", slice)
	}
}
