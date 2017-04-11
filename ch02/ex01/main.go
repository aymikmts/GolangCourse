package main

import (
	"fmt"

	"./tempconv"
)

func main() {
	// AbsoluteZeroC
	fmt.Print("AbsoluteZeroC\n")
	fmt.Printf("%v, %v, %v\n",
		tempconv.AbsoluteZeroC,
		tempconv.CToF(tempconv.AbsoluteZeroC),
		tempconv.CToK(tempconv.AbsoluteZeroC))

	fmt.Print("\nFreezingC\n")
	fmt.Printf("%v, %v, %v\n",
		tempconv.FreezingC,
		tempconv.CToF(tempconv.FreezingC),
		tempconv.CToK(tempconv.FreezingC))

	fmt.Print("\nBoilingC\n")
	fmt.Printf("%v, %v, %v\n",
		tempconv.BoilingC,
		tempconv.CToF(tempconv.BoilingC),
		tempconv.CToK(tempconv.BoilingC))
}
