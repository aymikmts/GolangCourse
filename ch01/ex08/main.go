// texte page	p.19
// date			2017/03/20

// mainは引数に接頭辞http://がなければ追加し、URLある内容を表示します
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		var urlStr string
		if strings.HasPrefix(url, "http://") {
			urlStr = url
		} else {
			urlStr = "http://" + url
		}

		resp, err := http.Get(urlStr)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", urlStr, err)
			os.Exit(1)
		}
		fmt.Printf("%s", b)
	}
}
