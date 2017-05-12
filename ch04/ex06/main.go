// Ex06は[]byteスライス内で隣接しているUnicodeスペースをASCIIスペースへ圧縮します。
// スライスは標準入力より入力された数値から生成します。
// 実装は"dedup/ex06.go"
package main

import (
	//"GolangCourse/ch04/dedup

	"bufio"
	"bytes"
	"fmt"
	"os"

	"../dedup"
)

func main() {

	// 配列の作成
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		var byteS []byte
		byteS = []byte(input.Text())

		fmt.Printf("input : %q\n", bytes.Runes(byteS))

		// Ex06の実行
		byteS = dedup.DedupSpace(byteS)

		fmt.Printf("output: %q\n\n", bytes.Runes(byteS))
	}
}
