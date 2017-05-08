// Ex03はint配列を直接逆順に並び替えます。
// 配列の大きさは"6"に固定です。
// 配列は標準入力より入力された数値から生成します。
package main

import (
	//"GolangCourse/ch04/rev""

	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"../rev"
)

func main() {

	// 配列の作成
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		var array [6]int
		s := strings.Fields(input.Text())
		for i, v := range s {
			if i >= len(array) {
				break
			}

			val, err := strconv.Atoi(v)
			if err != nil {
				fmt.Fprintf(os.Stderr, "v = %v, Atoi error: %v\n", v, err)
				os.Exit(1)
			}
			array[i] = val
		}

		fmt.Printf("input : %v\n", array)

		// Ex03の実行
		rev.ReverseByPointer(&array)

		fmt.Printf("output: %v\n", array)
	}
}
