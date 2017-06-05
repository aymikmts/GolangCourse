// Ex04は、セットの要素を含むスライスを返すElemsメソッドを実行します。
package main

import (
	"fmt"

	"../intset"
)

func main() {

	var x intset.IntSet
	x.Add(1)
	x.Add(65)
	fmt.Println(x.String())
	fmt.Println(x.Elems())
}
