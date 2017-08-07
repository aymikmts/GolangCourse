package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

const (
	CAPACITY = 10
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("usage: ./ex04 [max goroutine]\n")
		os.Exit(1)
	}
	maxGroutines, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalf(err.Error())
	}

	in := make(chan int)
	out := make(chan int)

	go job(in, out, maxGroutines)

	in <- 1
	start := time.Now()
	n := <-out
	end := time.Now()
	fmt.Printf("Finish. n:%d\n", n)

	// 時間を計測
	diff := end.Sub(start)
	avg := (int)(diff.Nanoseconds()) / n
	fmt.Printf("[%d goroutines]\ntotal: %v average: %v ns\n", n, diff, avg)
}

var count int

func job(in <-chan int, out chan<- int, maxGoroutines int) {
	next := make(chan int, CAPACITY)

	count++
	if count == maxGoroutines {
		for m := range in {
			out <- m
		}
	}
	go job(next, out, maxGoroutines)
	for n := range in {
		//fmt.Printf("n: %d\n", n)
		if n%10000 == 0 {
			fmt.Printf("%d\n", n)
		}
		next <- (n + 1)
	}
}
