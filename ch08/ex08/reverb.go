package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // 例: 接続が切れた
			continue
		}
		go handleConn(conn) // 接続を並行して処理する
	}
}

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)

	called := make(chan string)
	go func() {
		for {
			select {
			case <-time.After(10 * time.Second):
				fmt.Println("10 seconds.")
				c.Close()
				return
			case <-called:
			}
		}
	}()

	for input.Scan() {
		text := input.Text()
		called <- text
		go echo(c, text, 1*time.Second)
	}
	// 注意: input.Err()からの潜在的なエラーを無視している
	c.Close()
}
