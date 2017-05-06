// Ex11は、符号記号を持つ浮動小数点数にカンマを挿入します。
// 実装は、"comma/commaEx11.go"
package main

import (
	"fmt"
	"os"

	"../comma"
	//"GolangCourse/ch03/comma"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("in:%v  out:%v\n", os.Args[i], comma.CommaSignedFloat(os.Args[i]))
	}
}
