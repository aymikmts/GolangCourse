// Ex18は、fetchの振る舞いを変えることなくdeferで書込み可能なファイルを閉じます。
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
)

// fetchはURLをダウンロードして、ローカルファイルの名前と長さを返します。
func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}

	defer func() {
		if closeErr := f.Close(); err == nil {
			err = closeErr
		}
	}()

	n, err = io.Copy(f, resp.Body)

	return local, n, err
}

func main() {
	fname, n, err := fetch(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(fname, n)
}
