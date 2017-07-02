package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
)

var done = make(chan struct{})

type query struct {
	url      string
	filename string
	n        int64
	err      error
}

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

//!+
// Fetch downloads the URL and returns the
// name and length of the local file.
//func fetch(url string) (filename string, n int64, err error) {
func fetch(url string) query {
	var ret query

	req, err := http.NewRequest("GET", url, nil)

	cancel := make(chan struct{})
	req.Cancel = cancel
	if cancelled() {
		close(cancel)
		return ret
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		ret.err = err
		return ret
	}
	defer resp.Body.Close()

	local := path.Base(resp.Request.URL.Path)
	if local == "/" || local == "." {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		ret.err = err
		return ret
	}
	n, err := io.Copy(f, resp.Body)
	if closeErr := f.Close(); err == nil {
		err = closeErr
	}
	ret.url = url
	ret.filename = local
	ret.n = n
	ret.err = err
	close(done)
	return ret
}

func mirroredQuery(list []string) query {
	resp := make(chan query, len(list))
	for _, url := range list {
		go func() { resp <- fetch(url) }()
	}
	return <-resp
}

//!-

func main() {
	urls := os.Args[1:]
	fmt.Printf("request:%v\n", urls)
	resp := mirroredQuery(urls)
	if resp.err != nil {
		log.Fatal(resp.err)
	}
	fmt.Printf("first response:%s\nfile: %s  size:%d\n", resp.url, resp.filename, resp.n)
}
