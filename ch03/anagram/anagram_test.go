package anagram

import "testing"

func TestIsAnagram(t *testing.T) {
	var tests = []struct {
		input1 string
		input2 string
		want   bool
	}{
		{"hello, world!", "horlo, !welld", true},
		{"Hello, World!", "hoRlo, !welLd", true},
		{"こんにちは、世界", "世界、こんにちは", true},
		{"test", "foobar", false},
	}

	for _, test := range tests {
		got := IsAnagram(test.input1, test.input2)
		if got != test.want {
			t.Errorf("IsAnagram(%v, %v) = %v", test.input1, test.input2, got)
		}
	}
}

func TestSortString(t *testing.T) {
	var tests = []struct {
		input string
		want  string
	}{
		{"cdba", "abcd"},
		{"bCAd", "ACbd"},
		{"えういおあ", "あいうえお"},
	}

	for _, test := range tests {
		got := SortString(test.input)
		if got != test.want {
			t.Errorf("SortString(%v) = %v", test.input, got)
		}
	}

}
