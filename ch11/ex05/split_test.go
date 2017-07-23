package main

import (
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {
	tests := []struct {
		s     string
		sep   string
		words int
	}{
		{"", "", 0},
		{"a:b:c", "", 5},
		{"a:b:c", ":", 3},
		{"a::b::c", "::", 3},
	}
	for _, test := range tests {
		got := len(strings.Split(test.s, test.sep))
		if got != test.words {
			t.Errorf("Split(%q, %q) returned %d words, want %d", test.s, test.sep, got, test.words)
		}
	}
}
