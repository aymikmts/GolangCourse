// Ex01は、ワードと行に対するカウンタを実装します。
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

type LineCounter int
type WordCounter int

func (c *LineCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))

	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		*c++
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "scanner error: %v\n", err)
	}

	return len(p), nil
}

func (c *WordCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))

	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		*c++
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "scanner error: %v\n", err)
	}

	return len(p), nil
}

func main() {

	var wc WordCounter
	str := "hello world!\n Hello World!\n こんにちは　世界\n"
	fmt.Printf("str:\n %s\n", str)
	wc.Write([]byte(str))
	fmt.Printf("word count: %d\n", wc)

	var lc LineCounter
	lc.Write([]byte(str))
	fmt.Printf("line count: %d\n", lc)

}
