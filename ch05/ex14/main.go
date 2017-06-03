// Ex14は、url, topoSort, ファイル構造を表示します。
// フラグでどれを表示するかを切り替えます。
package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"gopl.io/ch5/links"
)

var out io.Writer = os.Stdout
var uFlag = flag.Bool("u", false, "crawl url")
var tFlag = flag.Bool("t", false, "show dependency of prereqs")
var dFlag = flag.Bool("d", false, "show directories tree")

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"liner algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

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
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

// crawlは、URLを表示し、リンクを抽出し、そして抽出されたリンクも訪れてリンクを返します。
func crawl(url string) []string {
	fmt.Fprintln(out, url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

// nodeDependenceは、topoSortの依存関係を返します。
func nodeDependence(node string) []string {
	fmt.Fprintln(out, node)
	var order []string
	order = prereqs[node]
	return order
}

// treePathは、ファイル構造を返します。
func treePath(path string) []string {
	fmt.Fprintln(out, path)
	var list []string

	f, err := os.Open(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot open %s\n", path)
		return nil
	}
	defer f.Close()

	fileInfo, err := f.Stat()
	if fileInfo.IsDir() {
		files, err := ioutil.ReadDir(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cannot read dir %s\n", fileInfo.Name())
			return nil
		}
		for _, file := range files {
			list = append(list, treePath(filepath.Join(path, file.Name()))...)
		}
	}
	return list
}

func main() {
	flag.Parse()

	if *uFlag {
		// コマンドライン引数から開始して、
		// ウェブを幅優先でクロールする
		breadthFirst(crawl, flag.Args())
	}

	if *tFlag {
		// topoSortの依存関係を表示する
		breadthFirst(nodeDependence, flag.Args())
	}

	if *dFlag {
		// ファイル構造を表示する
		breadthFirst(treePath, flag.Args())
	}
}
