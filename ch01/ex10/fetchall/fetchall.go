// fetchallは複数のURLの内容を並行して取り出します。
package fetchall

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func Fetch(inurl string, ch chan<- string) {
	start := time.Now()
	url := addprefix(inurl)
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // chチャネルへ送信
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // 資源をリークさせない
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}

func addprefix(url string) string {
	var ret string
	if strings.HasPrefix(url, "http://") {
		ret = url
	} else {
		ret = "http://" + url
	}
	return ret
}
