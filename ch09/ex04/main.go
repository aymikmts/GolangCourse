package main

import "fmt"

func main() {
	in := make(chan int)
	out := make(chan int)

	go job(in, out)

	in <- 1
	n := <-out
	fmt.Printf("Finish. n:%d\n", n)
}

var count int

func job(in <-chan int, out chan<- int) {
	next := make(chan int, 10)

	count++
	if count == 100 {
		for m := range in {
			out <- m
		}
	}
	go job(next, out)
	for n := range in {
		fmt.Printf("n: %d\n", n)
		next <- (n + 1)
	}
}
