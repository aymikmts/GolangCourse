package main

import (
	"bytes"
	"fmt"
	"os"
	"reflect"
	"strings"
	"testing"

	"golang.org/x/net/html"
)

func TestEachNode(t *testing.T) {
	var tests = []struct {
		input string
	}{
		{"testdata/input.html"},
		//	{"testdata/golang.org.html"},
	}
	for _, test := range tests {
		input, err := os.Open(test.input)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cannot open file: %s\n", test.input)
		}
		defer input.Close()

		inDoc, err := html.Parse(input)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to html.Parse.\n")
		}

		out = new(bytes.Buffer)
		forEachNode(inDoc, startElement, endElement)
		got := out.(*bytes.Buffer)

		outDoc, _ := html.Parse(got)
		if !isEqualData(inDoc, outDoc) {
			t.Errorf("Not equal between input and output.\n")
		}
	}
}

func isEqualData(in, out *html.Node) (ret bool) {
	ret = true

	if in.Type != out.Type {
		ret = false
		fmt.Fprintf(os.Stderr, "not match Type. %v != %v", in.Type, out.Type)
		return
	}

	inData := strings.TrimSpace(in.Data)
	outData := strings.TrimSpace(out.Data)
	if inData != outData {
		ret = false
		fmt.Fprintf(os.Stderr, "not match Data. %s != %s", inData, outData)
		return
	}
	if !reflect.DeepEqual(in.Attr, out.Attr) {
		ret = false
		fmt.Fprintf(os.Stderr, "not match Attribute.")
		return
	}

	for inc, outc := in.FirstChild, out.FirstChild; inc != nil && outc != nil; inc, outc = inc.NextSibling, outc.NextSibling {
		ret = isEqualData(inc, outc)
		if !ret {
			return
		}
	}

	return
}
