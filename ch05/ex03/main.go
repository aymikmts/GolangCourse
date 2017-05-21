// Ex03 は標準入力から読み込まれたHTMLドキュメント内のすべてのテキストノードの内容を表示します。。
package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ex03: %v\n", err)
		os.Exit(1)
	}
	printTextNode(nil, doc)
}

// printTextNodeは、HTMLドキュメントツリー内でテキストノードの内容を表示します。
func printTextNode(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data) // push tag
	} else if n.Type == html.TextNode {
		length := len(stack)
		if length > 0 {
			if stack[length-1] != "script" && stack[length-1] != "style" {
				if data := strings.TrimSpace(n.Data); len(data) > 0 {
					fmt.Printf("%s\n", data)
				}
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		printTextNode(stack, c)
	}
}
