package main

import (
	"GolangCourse/ch12/ex01/display"
	"fmt"
)

type A struct {
	id   int
	name string
}

type B struct {
	id      int
	name    string
	structA A
}

func main() {

	fmt.Println("--- key is struct) ---")
	mapA := make(map[A]bool)
	a1 := A{id: 1, name: "a1"}
	a2 := A{id: 2, name: "a2"}
	a3 := A{id: 3, name: "a3"}
	mapA[a1] = true
	mapA[a2] = false
	mapA[a3] = true
	display.Display("mapA", mapA)

	fmt.Println("\n--- key is nesting struct) ---")
	mapB := make(map[B]bool)
	b1 := B{id: 10, name: "b1", structA: a1}
	b2 := B{id: 20, name: "b2", structA: a2}
	b3 := B{id: 30, name: "b3", structA: a3}
	mapB[b1] = true
	mapB[b2] = false
	mapB[b3] = true
	display.Display("mapB", mapB)

	fmt.Println("\n--- key is array) ---")
	mapC := make(map[[3]int]bool)
	mapC[[3]int{1, 2, 3}] = true
	mapC[[3]int{10, 20, 30}] = false
	display.Display("mapC", mapC)

	fmt.Println("\n--- key is struct array) ---")
	mapD := make(map[[3]B]bool)
	mapD[[3]B{b1, b2, b3}] = true
	display.Display("mapD", mapD)
}
