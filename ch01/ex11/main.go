// mainは
package main

import (
	"fmt"
	"os"
	"time"

	"../ex10/fetchall"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetchall.Fetch(url, ch) // ゴルーチンを開始
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // chチャネルから受信
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}
