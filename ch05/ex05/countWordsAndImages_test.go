package main

import (
	"fmt"
	"io"
	"os"
	"testing"

	"golang.org/x/net/html"
)

func TestCountWordsAndImages(t *testing.T) {

	var tests = []struct {
		testData string
		words    int
		images   int
	}{
		{"test.html", 10, 2},
	}

	for _, test := range tests {
		n := readTestFile(test.testData)
		w, i := countWordsAndImages(nil, n)
		if w != test.words {
			t.Errorf("words is %d, but got is %d.", test.words, w)
		}
		if i != test.images {
			t.Errorf("images is %d, but got is %d.", test.images, i)
		}
	}
}

// test用HTMLファイルを読み込む
func readTestFile(fname string) *html.Node {
	file, err := os.Open(fname)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot open %s, %v\n", fname, err)
	}
	input := io.Reader(file)

	doc, err := html.Parse(input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
	return doc
}
