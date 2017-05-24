// Ex05 は
package main

import (
	"fmt"
	"os"
	"strings"

	"bufio"

	"golang.org/x/net/html"
)

func main() {
	// input := bufio.NewScanner(os.Stdin)
	// for input.Scan() {
	// 	url := input.Text()
	url := ""
	words, images, err := CountWordsAndImages(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ex05: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("words: %d\timages:%d\n", words, images)
	//}
}

// CountWordsAndImagesはHTMLドキュメントに対するHTTP GETリクエストをurlへ
// 行い、そのドキュメント内に含まれる単語と画像の数を返します。
func CountWordsAndImages(url string) (words, images int, err error) {
	// resp, err := http.Get(url)
	// if err != nil {
	// 	return
	// }
	// doc, err := html.Parse(resp.Body)
	// resp.Body.Close()
	// if err != nil {
	// 	err = fmt.Errorf("parsing HTML: %s", err)
	// 	return
	// }

	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ex05: %v\n", err)
		os.Exit(1)
	}
	words, images = countWordsAndImages(doc)
	return
}

// countWordsAndImagesは単語と画像の数を返します。
func countWordsAndImages(n *html.Node) (words, images int) {
	var stack []string
	switch n.Type {
	case html.ElementNode:
		stack = append(stack, n.Data) // push tag
		if stack[len(stack)-1] == "img" {
			images++
		}
	case html.TextNode:
		length := len(stack)
		if length > 0 {
			if stack[length-1] != "script" && stack[length-1] != "style" {
				scanner := bufio.NewScanner(strings.NewReader(n.Data))
				words += countWords(scanner)
			}
		}

	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		w, i := countWordsAndImages(c)
		words += w
		images += i
	}
	return
}

func countWords(input *bufio.Scanner) int {
	var counts int
	input.Split(bufio.ScanWords)
	for input.Scan() {
		counts++
	}
	return counts
}
