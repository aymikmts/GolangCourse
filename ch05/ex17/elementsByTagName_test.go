package main

import (
	"fmt"
	"os"
	"testing"

	"golang.org/x/net/html"
)

func TestElementsByTagName(t *testing.T) {
	var tests = []struct {
		fname string
		name  []string
		num   []int
	}{
		{"testdata/input.html", []string{}, []int{}},
		{"testdata/input.html", []string{"h2"}, []int{0}},
		{"testdata/input.html", []string{"img", "h1"}, []int{2, 3}},
	}

	for _, test := range tests {
		f, err := os.Open(test.fname)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cannot open: %s err:%s\n", test.fname, err)
			continue
		}
		defer f.Close()

		doc, err := html.Parse(f)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		m := make(map[string]int)
		got := ElementsByTagName(doc, test.name...)
		for _, n := range got {
			m[n.Data]++
		}

		for i := 0; i < len(m); i++ {
			if m[test.name[i]] != test.num[i] {
				t.Errorf("[%s]num is %d, but got is %d\n", test.name[i], test.num[i], m[test.name[i]])
			}
		}
	}
}
