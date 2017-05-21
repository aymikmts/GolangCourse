// Ex02 は標準入力から読み込まれたHTMLドキュメント内の要素の数を表示します。
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ex02: %v\n", err)
		os.Exit(1)
	}

	m := make(map[string]int)
	mapElementCount(&m, doc)

	fmt.Printf("Element\t\tCount\n")
	for key, val := range m {
		fmt.Printf("%-10s\t%d\n", key, val)
	}
}

// mapElementCountは、HTMLドキュメントツリー内でその要素名を持つ要素の数を対応させるmapを作成します。
func mapElementCount(m *map[string]int, n *html.Node) {

	for n == nil {
		return
	}

	if n.Type == html.ElementNode {
		(*m)[n.Data]++
	}

	mapElementCount(m, n.FirstChild)
	mapElementCount(m, n.NextSibling)
}
