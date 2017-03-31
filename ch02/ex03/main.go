// mainはpopcountの出力を行います。

package main

import (
	"fmt"

	"./popcount"
	"./popcountroop"
)

func main() {
	var input uint64
	input = 255
	output := popcount.PopCount(255)
	fmt.Printf("nonroop: [in]%b [out]%d\n", input, output)

	output = popcountroop.PopCount(255)
	fmt.Printf("roop:    [in]%b [out]%d\n", input, output)

}
