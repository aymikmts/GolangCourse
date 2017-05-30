package main

import (
	"bytes"
	"os"
	"testing"

	"golang.org/x/net/html"
)

func TestByNamelessFunc(t *testing.T) {
	var tests = []struct {
		input string
	}{
		{"testdata/test.html"},
		{"testdata/input.html"},
		{"testdata/golang.org.html"},
		{"testdata/theta360.com.html"},
	}

	for _, test := range tests {
		input, err := os.Open(test.input)
		if err != nil {
			t.Fatalf("cannot open file: %s\n", test.input)
		}
		defer input.Close()

		doc, err := html.Parse(input)
		if err != nil {
			t.Fatalf("failed to html.Parse().\n")
		}

		out = new(bytes.Buffer)
		byNamelessFunc(doc)
		got := bytes.NewBuffer(out.(*bytes.Buffer).Bytes())

		out = new(bytes.Buffer)
		forEachNode(doc, startElement, endElement)

		if bytes.Compare(got.Bytes(), out.(*bytes.Buffer).Bytes()) != 0 {
			t.Errorf("want size is %d, but got size is %d.\n", got.Len(), out.(*bytes.Buffer).Len())
		}
	}
}
