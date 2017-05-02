//Ex10は、bytes.Bufferを用いて再帰呼び出しをせずに、10進表記整数文字列にカンマを挿入します。
package main

import (
	"fmt"
	"os"

	"../comma"
	//"GolangCourse/ch03/comma"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("in:%v  out:%v\n", os.Args[i], comma.CommaWithBuffer(os.Args[i]))
	}
}
