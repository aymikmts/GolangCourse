// Ex03は、IntersectWith, DifferenceWith, SymmetricDifferenceを実行します。
package main

import (
	"flag"
	"fmt"

	"../intset"
)

var union = flag.Bool("u", false, "do UnionWith()")
var intersect = flag.Bool("i", false, "do IntersectWith()")
var diff = flag.Bool("d", false, "do DifferenceWith()")
var symdiff = flag.Bool("s", false, "do SymmetricDifference()")

func main() {
	flag.Parse()

	var x, y intset.IntSet
	x.Add(1)
	x.Add(9)
	x.Add(65)
	x.Add(144)
	fmt.Println(x.String())

	y.Add(5)
	y.Add(65)
	y.Add(144)
	fmt.Println(y.String())

	if *union {
		x.UnionWith(&y)
	}
	if *intersect {
		x.IntersectWith(&y)
	}
	if *diff {
		x.DifferenceWith(&y)
	}
	if *symdiff {
		x.SymmetricDifference(&y)
	}
	fmt.Println(x.String())
}
