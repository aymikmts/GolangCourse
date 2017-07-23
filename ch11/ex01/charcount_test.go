package main

import (
	"strings"
	"testing"
)

func TestCharCount(t *testing.T) {
	var tests = []struct {
		input   string
		counts  map[rune]int
		utflen  []int
		invalid int
	}{
		{"", map[rune]int{}, []int{0, 0, 0, 0, 0}, 0},
		{"test", map[rune]int{'t': 2, 'e': 1, 's': 1}, []int{0, 4, 0, 0, 0}, 0},
		{"ああ　いいい \naaiii", map[rune]int{'あ': 2, 'い': 3, 'a': 2, 'i': 3, ' ': 1, '\n': 1, '　': 1}, []int{0, 7, 0, 6, 0}, 0},
	}

	for _, test := range tests {
		r := strings.NewReader(test.input)
		counts, utflen, invalid, err := charcount(r)
		if err != nil {
			t.Errorf("unexpected err: %v", err)
			continue
		}
		for c, n := range counts {
			if n != test.counts[c] {
				t.Errorf("%q count = %d, want %d", c, n, test.counts[c])
				continue
			}
		}
		for i, n := range utflen {
			if n != test.utflen[i] {
				t.Errorf("utflen %d = %d, want %d", i, n, test.utflen[i])
				continue
			}
		}
		if invalid != test.invalid {
			t.Errorf("invalid %d, want %d", invalid, test.invalid)
		}
	}
}
