// Ex01は、ビットベクタのLen, Remove, Clear, Copyを実行します。
package main

import (
	"fmt"

	"./intset"
)

func main() {
	var x intset.IntSet
	fmt.Printf("Add: 1, 144, 9, 10\n")
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(10)
	fmt.Printf("%s\nlen: %d\n", x.String(), x.Len())

	fmt.Printf("\nRemove: 10, 2\n")
	x.Remove(10)
	x.Remove(2)
	fmt.Printf("%s\nlen: %d\n", x.String(), x.Len())

	fmt.Printf("\nClear:\n")
	x.Clear()
	fmt.Printf("%s\nlen: %d\n", x.String(), x.Len())

	var src intset.IntSet
	var dst *intset.IntSet
	src.Add(1)
	src.Add(10)
	src.Add(100)
	fmt.Printf("Copy:\n")
	fmt.Printf("\nsrc: %s\n", src.String())
	dst = src.Copy()
	fmt.Printf("copy src: %s\n", dst.String())

}
