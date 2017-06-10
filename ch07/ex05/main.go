// Ex05はLimitReaderを実行します。
package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
)

type limitReader struct {
	r io.Reader
	n int64
}

func (lr *limitReader) Read(b []byte) (n int, err error) {
	if lr.n <= 0 {
		return 0, io.EOF
	}
	if int64(len(b)) > lr.n {
		b = b[0:lr.n]
	}
	n, err = lr.r.Read(b)
	if err != nil {
		return 0, err
	}
	lr.n -= int64(n)
	return n, nil
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &limitReader{r, n}
}

func main() {
	text := []byte("Hello, World!\nこんにちは、世界\n")
	fmt.Printf("input:\n%s\n", string(text))
	reader := bytes.NewReader(text)

	lmt := LimitReader(reader, 6)
	buf, err := ioutil.ReadAll(lmt)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("output:\n%s\n", string(buf))
}
