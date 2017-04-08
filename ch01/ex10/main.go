// mainは
package main

import (
	"fmt"
	"os"
	"time"

	"./fetchall"
)

func main() {
	fetch()
	fetch()
}

func fetch() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetchall.Fetch(url, ch) // ゴルーチンを開始
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // chチャネルから受信
	}
	fmt.Fprintf(os.Stdout, "%.2fs elapsed\n", time.Since(start).Seconds())
}
