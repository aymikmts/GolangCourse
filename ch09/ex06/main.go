package main

import (
	"sync"
)

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
func main() {
	const n = 35
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			_ = fib(n)
			//fmt.Printf("[%2d]Fibonacci(%d) = %d\n", i, n, fibN)
		}(i)
	}
	wg.Wait()
	//fmt.Printf("done.\n")
}
