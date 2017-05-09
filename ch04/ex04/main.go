// Ex04はintスライスを回転させます。
// スライスは標準入力より入力された数値から生成します。
// 標準入力の1つ目の値はRotateする要素数で、それ以降の数値からスライスを生成します。
// 要素数に正の値を入れると左回転、負の値を入れると右回転します。
// 例) "2 10 20 30 40 50" → スライス{10, 20, 30, 40, 50}を左へ2要素分ずらす。
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
	var rot int

	// 配列の作成
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		var slice []int
		s := strings.Fields(input.Text())
		for i, v := range s {
			val, err := strconv.Atoi(v)
			if err != nil {
				fmt.Fprintf(os.Stderr, "v = %v, Atoi error: %v\n", v, err)
				os.Exit(1)
			}

			// 一番はじめは回転数
			if i == 0 {
				rot = val
				continue
			}
			slice = append(slice, val)
		}

		fmt.Printf("input : %v\nrot: %d\n", slice, rot)

		// Ex04の実行
		rev.Rotate(slice, rot)

		fmt.Printf("output: %v\n\n", slice)
	}
}
