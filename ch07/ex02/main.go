// Ex02は、CountingWriterを実装します。
package main

import (
	"fmt"
	"io"
	"os"
)

type NewWriter struct {
	w io.Writer
	c int64
}

func (w *NewWriter) Write(p []byte) (int, error) {
	n, err := w.w.Write(p)
	w.c += int64(n)
	return n, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	var writer NewWriter
	writer.w = w

	return &writer, &writer.c
}

func main() {
	str := "Hello, world!"

	w, cnt := CountingWriter(os.Stdout)
	fmt.Printf("type:%T, WriteCount:%d\n", w, *cnt)

	fmt.Println()

	fmt.Fprintf(w, "%s\n", str)
	fmt.Printf("type:%T, WriteCount:%d\n", w, *cnt)

	fmt.Println()

	fmt.Fprintf(w, "test\n")
	fmt.Printf("type:%T, WriteCount:%d\n", w, *cnt)
}
