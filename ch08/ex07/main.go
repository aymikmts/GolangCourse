package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"

	"gopl.io/ch5/links"
)

var targetHost *url.URL

//var targetHost string
var out io.Writer = os.Stdout

// breadthFirstはworklist内の個々の項目に対してfを呼び出します。
// fから返されたすべての項目はworklistへ追加されます。
// fは、それぞれの項目に対して高々一度しか呼び出されません。
func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				if isTargetURL(item) {
					go download(item)
					worklist = append(worklist, f(item)...)
				}
			}
		}
	}
}

func isTargetURL(item string) bool {
	u, err := url.Parse(item)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return false
	}

	return strings.HasSuffix(u.Host, targetHost.Host)
}

//
func download(item string) {
	resp, err := http.Get(item)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	dir := path.Dir(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}

	if strings.HasSuffix(item, "/") {
		if strings.HasSuffix(dir, local) {
			local = "index.html"
		}
	}

	dirName := "data/" + targetHost.Host + dir
	err = os.MkdirAll(dirName, 0777)
	if err != nil {
		log.Fatalln(err)
	}

	f, err := os.Create(dirName + "/" + local)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	fmt.Printf("%s\n", f.Name())

	_, err = io.Copy(f, resp.Body)
	if err != nil {
		log.Println(err)
	}
}

// crawlは、URLを表示し、リンクを抽出し、抽出されたリンクも訪れるようにリンクを返します。
func crawl(url string) []string {
	//fmt.Fprintln(out, url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {

	u, err := url.Parse(os.Args[1])
	if err != nil {
		log.Printf("failed to parse url: %s err: %v\n", os.Args[1], err)
	}

	targetHost = u
	//tgargetHost = r/" + u.Host
	//fmt.Fprintf(out, "targetHost: %v\n", targetHost)
	// コマンドライン引数から開始して、
	// ウェブを幅優先でクロールする
	breadthFirst(crawl, os.Args[1:])
}
