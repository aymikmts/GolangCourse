// Ex02は、可変長個数に対応したAddAllを実行します。
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"../intset"
)

func main() {
	var x intset.IntSet
	var input []int
	for _, s := range os.Args[1:] {
		val, err := strconv.Atoi(s)
		if err != nil {
			log.Fatalln(err)
		}
		input = append(input, val)
	}

	x.AddAll(input...)
	fmt.Println(x.String())
}
