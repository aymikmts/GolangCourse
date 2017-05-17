// Ex07はUTF-8でエンコードされた文字列を表す[]byteスライスを直接逆順に並び替えます。
// スライスは標準入力より入力された数値から生成します。
// 実装は"rev/ex07.go"
package main

import (
	//"GolangCourse/ch04/rev"

	"bufio"
	"fmt"
	"os"

	"../rev"
)

func main() {

	// 配列の作成
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		var array []byte

		array = []byte(input.Text())

		fmt.Printf("input:\n %[1]s\n", array)

		// Ex07の実行
		rev.ReverseUnicodeByte(array)

		fmt.Printf("output:\n %[1]s\n\n", array)
	}
}
