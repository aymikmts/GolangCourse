// mainは、個々の引数のインデックスと値の組を1行ごとに出力します。
package main

import (
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Println(i, os.Args[i])
	}
}
