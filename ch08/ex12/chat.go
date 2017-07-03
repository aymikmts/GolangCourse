package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type client struct {
	msg  chan<- string // 送信用メッセージチャネル
	user string
}

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // クライアントから受信するすべてのメッセージ
)

func breadcaster() {
	clients := make(map[client]bool) // すべての接続されているクライアント
	for {
		select {
		case msg := <-messages:
			// 受信するメッセージのすべてのクライアントの
			// 送信用メッセージチャネルへブロードキャストする
			for cli := range clients {
				cli.msg <- msg
			}

		case cli := <-entering:
			clients[cli] = true
			for c := range clients {
				cli.msg <- "USER: " + c.user
			}

		case cli := <-leaving:
			delete(clients, cli)
			close(cli.msg)
		}
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string) // 送信用のクライアントメッセージ
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- client{ch, who}

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
	}
	// 注意: input.Err()からの潜在的なエラーを無視している

	leaving <- client{ch, who}
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // 注意: ネットワークのエラーを無視している
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go breadcaster()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
			continue
		}
		go handleConn(conn)
	}
}
