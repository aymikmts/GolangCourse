// Ex04は[]intを
// スライスは標準入力より入力された数値から生成します。
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

		fmt.Printf("input : %v\n", slice)

		// Ex04の実行
		rev.Rotate(slice, rot)

		fmt.Printf("output: %v\n", slice)
	}
}
