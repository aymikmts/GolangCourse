// Ex12は、outlineのpre、post関数を無名関数にして実装します。
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
		fmt.Fprintf(os.Stderr, "Ex12: %v\n", err)
		os.Exit(1)
	}

	fmt.Fprintf(out, "[usingSeparateFunc]\n")
	forEachNode(doc, startElement, endElement)

	fmt.Fprintf(out, "\n[byNamelessFunc]\n")
	byNamelessFunc(doc)
}

func byNamelessFunc(n *html.Node) {
	var depth int
	forEachNode(n,
		func(n *html.Node) {
			if n.Type == html.ElementNode {
				fmt.Fprintf(out, "%*s<%s>\n", depth*2, "", n.Data)
				depth++
			}
		},
		func(n *html.Node) {
			if n.Type == html.ElementNode {
				depth--
				fmt.Fprintf(out, "%*s</%s>\n", depth*2, "", n.Data)
			}
		})
}

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

var Depth int

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Fprintf(out, "%*s<%s>\n", Depth*2, "", n.Data)
		Depth++
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		Depth--
		fmt.Fprintf(out, "%*s</%s>\n", Depth*2, "", n.Data)
	}
}
