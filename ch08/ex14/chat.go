package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

type client struct {
	msg  chan<- string // 送信用メッセージチャネル
	addr string        // クライアントのネットワークアドレス
	name string        // ユーザーネーム
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
				cli.msg <- "USER: " + c.name
			}

		case cli := <-leaving:
			delete(clients, cli)
			close(cli.msg)
		}
	}
}

func getUserName(conn net.Conn) string {
	var name string
	fmt.Fprintf(conn, "please enter your name:\n")
	input := bufio.NewScanner(conn)
	for input.Scan() {
		name = input.Text()
		break
	}
	return name
}

func handleConn(conn net.Conn, name string) {
	ch := make(chan string) // 送信用のクライアントメッセージ
	go clientWriter(conn, ch)

	addr := conn.RemoteAddr().String()
	ch <- "You are " + name
	messages <- name + " has arrived"
	entering <- client{ch, addr, name}

	msg := make(chan string)
	go func() {
		for {
			select {
			case s := <-msg:
				messages <- s
			case <-time.After(5 * time.Minute):
				conn.Close()
			}
		}
	}()
	input := bufio.NewScanner(conn)
	for input.Scan() {
		msg <- name + ": " + input.Text()
	}
	// 注意: input.Err()からの潜在的なエラーを無視している

	leaving <- client{ch, addr, name}
	messages <- name + " has left"
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
		name := getUserName(conn)
		go handleConn(conn, name)
	}
}
