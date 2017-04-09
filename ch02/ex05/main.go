package main

import (
	"fmt"

	"./popcountandoperation"
)

func main() {
	val := popcountandoperation.PopCount(255)
	fmt.Println(val)
	val = popcountandoperation.PopCount(254)
	fmt.Println(val)
	val = popcountandoperation.PopCount(0)
	fmt.Println(val)
}
