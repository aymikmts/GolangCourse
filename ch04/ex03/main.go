// Ex03はint配列を直接逆順に並び替えます。
// 配列の大きさは"6"に固定です。
// 配列はコマンド引数より生成します。
package main

import (
	//"GolangCourse/ch04/rev""

	"fmt"
	"os"
	"strconv"

	"../rev"
)

func main() {
	// 配列の作成
	var array [6]int
	for i, v := range os.Args {
		if i == 0 {
			continue
		}
		if i > len(array) {
			break
		}

		val, err := strconv.Atoi(v)
		if err != nil {
			fmt.Fprintf(os.Stderr, "v = %v, Atoi error: %v\n", v, err)
			os.Exit(1)
		}
		array[i-1] = val
	}

	fmt.Printf("input : %v\n", array)

	rev.ReverseByPointer(&array)
	fmt.Printf("output: %v\n", array)
}
