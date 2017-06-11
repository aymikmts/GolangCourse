package main

import "testing"

func TestIsPalindrome(t *testing.T) {
	var tests = []struct {
		input string
		want  bool
	}{
		{"", true},
		{"a", true},
		{"aa", true},
		{"ab", false},
		{"kayak", true},
		{"こんにちは", false},
		{"たけやぶやけた", true},
	}
	for _, test := range tests {
		got := IsPalindrome(runes(test.input))
		if got != test.want {
			t.Errorf("IsPalindrome(%v) = %v", test.input, got)
		}
	}
}
