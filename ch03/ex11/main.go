// Ex11は、符号機能を持つ浮動小数点数表記にします。
// 実装は、"comma/commaEx11.go"
// 入力には符号を含んだ(あるいは含まない)数値文字列が入ることを前提としており、想定外の値が入力されたときのエラー処理は行なっていない。
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
