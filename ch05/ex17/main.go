// Ex17は、HTMLノードツリーと0個以上の名前が与えられたら、それらの名前と一致する要素すべてを返します。
package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/net/html"
)

func ElementsByTagName(doc *html.Node, name ...string) []*html.Node {
	if len(name) < 1 {
		return nil
	}

	var node []*html.Node

	var visit func(n *html.Node, tag ...string)

	visit = func(n *html.Node, tag ...string) {
		if n.Type == html.ElementNode {
			for _, t := range tag[:] {
				if t == n.Data {
					node = append(node, n)
				}
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			visit(c, tag...)
		}
	}

	visit(doc, name...)
	return node
}

func main() {

	file, err := os.Open(os.Args[len(os.Args)-1])
	if err != nil {
		log.Fatalf("failed to open file: %s\n", os.Args[len(os.Args)-1])
	}
	defer file.Close()

	doc, err := html.Parse(file)
	if err != nil {
		log.Fatalln(err)
	}

	node := ElementsByTagName(doc, os.Args[1:]...)
	for _, n := range node {
		fmt.Println(n.Data, n.Attr)
	}

}
