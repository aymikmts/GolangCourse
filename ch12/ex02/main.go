package main

import (
	"GolangCourse/ch12/ex02/display"
	"fmt"
)

func main() {
	type Cycle struct {
		Value int
		Tail  *Cycle
	}
	var c Cycle
	c = Cycle{42, &c}

	fmt.Println("--- Depth = 3 ---")
	display.Display("c", c)

	fmt.Println("\n--- Depth = 10 ---")
	display.Depth = 10
	display.Display("c", c)
}
