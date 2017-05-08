// Ex03はint配列を直接逆順に並び替えます。
// 配列の大きさは"6"に固定です。
package main

import (
	//"GolangCourse/ch04/rev""

	"fmt"

	"../rev"
)

func main() {
	a := [6]int{0, 1, 2, 3, 4, 5}
	rev.ReverseByPointer(&a)
	fmt.Println(a)

	b := [6]int{1, 2, 3}
	rev.ReverseByPointer(&b)
	fmt.Println(b)
}
