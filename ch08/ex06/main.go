package main

import (
	"flag"
	"fmt"
	"log"

	"gopl.io/ch5/links"
)

type workinput struct {
	lists []string
	depth int
}

var d = flag.Int("depth", 3, "set crawl depth.")
var tokens = make(chan struct{}, 20)

func crawl(url string, depth int) *workinput {
	if depth > *d {
		return &workinput{nil, depth + 1}
	}
	fmt.Println(depth, url)
	tokens <- struct{}{} // トークンを獲得
	list, err := links.Extract(url)
	<-tokens // トークンを開放
	if err != nil {
		log.Print(err)
	}
	return &workinput{list, depth + 1}
}

func main() {
	flag.Parse()

	worklist := make(chan *workinput)
	var n int // worklistへの送信待ちの数

	// コマンドラインの引数で開始する
	n++
	go func() { worklist <- &workinput{flag.Args(), 0} }()

	// ウェブを平行にクロールする
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		//depth := list.depth
		for _, link := range list.lists {
			if !seen[link] {
				seen[link] = true
				n++
				depth := list.depth
				go func(link string, depth int) {
					worklist <- crawl(link, depth)
				}(link, depth)
			}
		}
	}
}
