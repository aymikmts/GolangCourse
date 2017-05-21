// Ex01 は標準入力から読み込まれたHTMLドキュメント内のリンクを表示します
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visitByRecursive(nil, doc) {
		fmt.Println(link)
	}
}

// visitByRecursiveは、再帰を使ってn内で見つかったリンクを一つひとつlinksへ追加し、その結果を返します。
func visitByRecursive(links []string, n *html.Node) []string {

	for n == nil {
		return links
	}

	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	links = visitByRecursive(links, n.FirstChild)
	links = visitByRecursive(links, n.NextSibling)

	return links
}
