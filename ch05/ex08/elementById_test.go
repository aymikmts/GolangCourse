package main

import (
	"fmt"
	"os"
	"testing"

	"golang.org/x/net/html"
)

func TestEachNode(t *testing.T) {
	var tests = []struct {
		searchId string
		input    string
		hit      bool
		wantData string
		wantID   string
	}{
		{"test", "testdata/input.html", true, "h1", "test"},
		{"id1", "testdata/test1.html", true, "p", "id1"}, // 同idが複数あるうちの最初のものを選ぶかどうか
		{"id2", "testdata/test2.html", true, "a", "id2"}, // idが複数あるうちの最初のものを選ぶかどうか
		{"noid", "testdata/test2.html", false, "", ""},   // idがない場合
	}
	for _, test := range tests {
		input, err := os.Open(test.input)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cannot open file: %s\n", test.input)
		}
		defer input.Close()

		doc, err := html.Parse(input)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to html.Parse.\n")
		}

		node := ElementByID(doc, test.searchId)
		if node != nil {
			if !test.hit {
				t.Errorf("\"%s\" has been wanted not to hit. but has hit.\n", test.searchId)
			}
			if node.Type != html.ElementNode {
				t.Errorf("node type is %d, but want is %d.\n", node.Type, html.ElementNode)
			}
			if node.Data != test.wantData {
				t.Errorf("node data is %s, but want is %s.\n", node.Data, test.wantData)
			}
			for _, a := range node.Attr {
				if a.Key == "id" {
					if a.Val != test.wantID {
						t.Errorf("node id is %s, but want is %s.\n", a.Val, test.wantID)
					}
				}
			}
		} else {
			if test.hit {
				t.Errorf("\"%s\" has been wanted to hit, but has not hit.\n", test.searchId)
			}
		}
	}
}
