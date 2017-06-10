// Ex04は、HTMLパーサが文字列からの入力を受け付けるようなNewReaderメソッドを実装します。
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"golang.org/x/net/html"
)

var out io.Writer = os.Stdout

// strings.goを参考にstructを作成
type strReader struct {
	s string
	i int64
}

// strings.goを参考にReadを作成
func (r *strReader) Read(p []byte) (n int, err error) {
	if r.i >= int64(len(r.s)) {
		return 0, io.EOF
	}
	n = copy(p, r.s[r.i:])
	r.i += int64(n)
	return n, err
}

func NewReader(s string) io.Reader {
	return &strReader{s, 0}
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}

func Parse(str string) {
	doc, err := html.Parse(NewReader(str))
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Fprintln(out, link)
	}
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open file: %s err: %v\n", os.Args[1], err)
		os.Exit(1)
	}
	defer file.Close()

	buf, err := ioutil.ReadAll(file)
	Parse(string(buf))
}
