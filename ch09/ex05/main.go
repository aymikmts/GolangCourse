package main

import (
	"fmt"
	"time"
)

var durationTime = time.Duration(1 * time.Second)

func playerA(rcv <-chan string, snd chan<- string, count <-chan int) {
	for {
		select {
		case msg := <-rcv:
			//fmt.Printf("A: %s\n", msg)
			snd <- msg
		case <-count:
		default:
		}
	}
}

func playerB(rcv <-chan string, snd chan<- string, count chan<- int) {
	var i int
	for {
		select {
		case msg := <-rcv:
			//fmt.Printf("B:[%d] %s\n", i, msg)
			i++
			count <- i
			snd <- msg
		default:
		}
	}
}

func main() {

	messageA := make(chan string)
	messageB := make(chan string)
	count := make(chan int)

	go playerA(messageA, messageB, count)
	go playerB(messageB, messageA, count)

	start := time.Now()
	tick := time.Tick(durationTime)
	messageA <- "test"
	<-tick
	end := time.Now().Sub(start)
	c := <-count

	fmt.Printf("time: %v, count: %d\n", end, c)
	fmt.Printf("%f per second\n", float64(c)/end.Seconds())
}
