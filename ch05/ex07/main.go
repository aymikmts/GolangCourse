// Ex07は、コメントノード、テキストノード、個々の要素ノードを表示するHTMLプリティプリンタです。
// 要素が子を持たない場合は、<img/>のような短い形式で出力します。
package main

import (
	"fmt"
	"io"
	"os"

	"golang.org/x/net/html"
)

var out io.Writer = os.Stdout

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ex07 error: %v\n", err)
	}

	forEachNode(doc, startElement, endElement)
}

// forEachNodeは、nから始まるツリー内の個々のノードxに対して
// 関数pre(x)とpost(x)を呼び出します。その二つの関数はオプションです。
// preは子ノードを訪れる前に呼び出され、
// postは子ノードを訪れた後に呼び出されます。
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

var depth int
var noChild bool

func startElement(n *html.Node) {
	switch n.Type {
	case html.TextNode:
		fmt.Fprintf(out, "%*s%s\n", depth*2, "", n.Data)
	case html.ElementNode:
		fmt.Fprintf(out, "%*s<%s", depth*2, "", n.Data)
		for _, a := range n.Attr {
			fmt.Fprintf(out, " %s='%s'", a.Key, a.Val)
		}

		if n.FirstChild == nil {
			fmt.Fprintf(out, "/>\n")
			noChild = true
			return
		}

		fmt.Fprintf(out, ">\n")
		depth++

	case html.CommentNode:
		fmt.Fprintf(out, "%*s<!-- %s -->\n", depth*2, "", n.Data)
	}
}

func endElement(n *html.Node) {
	if noChild {
		noChild = false
		return
	}

	if n.Type == html.ElementNode {
		depth--
		fmt.Fprintf(out, "%*s</%s>\n", depth*2, "", n.Data)
	}
}
