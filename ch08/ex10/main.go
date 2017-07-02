package main

import (
	"fmt"
	"log"
	"os"

	"GolangCourse/ch08/ex10/links"
)

var done = make(chan struct{})

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url, done)
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	// 入力を検出するとキャンセルを伝える
	go func() {
		os.Stdin.Read(make([]byte, 1)) // 1バイトを読み込む
		close(done)
		fmt.Println("CANCEL REQUEST")
	}()

	worklist := make(chan []string)  // URLのリスト、重複を含む
	unseenLinks := make(chan string) // 重複してないURL

	// コマンドラインの引数で開始する
	go func() { worklist <- os.Args[1:] }()

	// 未探索のリンクを取得するために20個のクローラのゴルーチンを生成する
	for i := 0; i < 20; i++ {
		go func() {
			select {
			case <-done:
				fmt.Printf("CANCELED\n")
				return
			default:
				for link := range unseenLinks {
					foundLinks := crawl(link)
					go func() { worklist <- foundLinks }()
				}
			}
		}()
	}

	// メインゴルーチンはworklistの項目の重複をなくし、
	// 未探索の項目をクローラへ送る
	seen := make(map[string]bool)
	for {
		select {
		case <-done:
			fmt.Printf("CANCELED\n")
			return
		case list := <-worklist:
			for _, link := range list {
				if !seen[link] {
					seen[link] = true
					unseenLinks <- link
				}
			}
		}
	}
}
