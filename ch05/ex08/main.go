// Ex08は、指定されたid属性を持つ最初のHTML要素を見つけます。
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"golang.org/x/net/html"
)

var out io.Writer = os.Stdout

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		input := strings.Fields(scanner.Text())

		file, err := os.Open(input[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to open file: %s\n", input[1])
		}
		defer file.Close()

		doc, err := html.Parse(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Ex08 error: %v\n", err)
		}

		fmt.Fprintf(out, "search ID:\n  %s\nresult:\n", input[0])
		node := ElementByID(doc, input[0])
		if node != nil {
			fmt.Fprintf(out, "  data: %s\n  attribute: %v\n", node.Data, node.Attr)
		} else {
			fmt.Fprintf(out, "  no hit.\n")
		}
	}
}

func forEachNode(n *html.Node, id string, pre, post func(n *html.Node, id string) bool) (bool, *html.Node) {

	if pre != nil {
		if pre(n, id) {
			return true, n
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		hit, node := forEachNode(c, id, pre, post)
		if hit {
			return true, node
		}
	}

	if post != nil {
		if post(n, id) {
			return true, n
		}
	}

	return false, nil
}

func preSearchElement(n *html.Node, id string) bool {
	if n.Type == html.ElementNode {
		for _, a := range n.Attr {
			if a.Key == "id" && a.Val == id {
				return true
			}
		}
	}

	return false
}

func postSearchElement(n *html.Node, id string) bool {
	if n.Type == html.ElementNode {
		for _, a := range n.Attr {
			if a.Key == "id" && a.Val == id {
				return true
			}
		}
	}

	return false
}

func ElementByID(doc *html.Node, id string) *html.Node {
	_, node := forEachNode(doc, id, preSearchElement, postSearchElement)
	return node
}
