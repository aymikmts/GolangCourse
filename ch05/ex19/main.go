// Ex19は、returnを使わずにpanicとrecoverで値を返します。
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func double(x int) (ret int) {
	defer func() {
		if p := recover(); p != nil {
			ret = x + x
		}
	}()

	panic(x)
}
func main() {
	val, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(double(val))
}
