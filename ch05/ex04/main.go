// Ex04 は標準入力から読み込まれたHTMLドキュメント内の、画像、スクリプト、スタイルシートなどのリンクを抽出します。
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ex04: %v\n", err)
		os.Exit(1)
	}

	links := make(map[string][]string)
	visitLinks(&links, doc)
	for key, vals := range links {
		fmt.Printf("[%v]\n", key)
		for _, val := range vals {
			fmt.Printf("  %v\n", val)
		}
	}
}

// visitLinksは、画像、スクリプト、スタイルシートなどの種類のリンクを抽出します。
func visitLinks(links *map[string][]string, n *html.Node) {
	if n.Type == html.ElementNode {
		for _, a := range n.Attr {
			var val string
			switch n.Data {
			case "a":
				if a.Key == "href" {
					val = a.Val
				}
			case "link":
				if a.Key == "href" {
					val = a.Val
				}
			case "img":
				if a.Key == "src" {
					val = a.Val
				}
			case "script":
				if a.Key == "src" {
					val = a.Val
				}
			default:
				continue
			}

			if val != "" {
				(*links)[n.Data] = append((*links)[n.Data], val)
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visitLinks(links, c)
	}

}
