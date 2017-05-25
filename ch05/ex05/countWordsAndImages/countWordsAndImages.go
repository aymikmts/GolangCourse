package countwordsandimages

import (
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/net/html"

	"bufio"
)

// CountWordsAndImagesはHTMLドキュメントに対するHTTP GETリクエストをurlへ
// 行い、そのドキュメント内に含まれる単語と画像の数を返します。
func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}

	words, images = countWordsAndImages(nil, doc)
	return
}

// countWordsAndImagesは単語と画像の数を返します。
func countWordsAndImages(stack []string, n *html.Node) (words, images int) {
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
				scanner.Split(bufio.ScanWords)
				for scanner.Scan() {
					words++
				}
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		w, i := countWordsAndImages(stack, c)
		words += w
		images += i
	}

	return
}
