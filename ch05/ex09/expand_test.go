package main

import "testing"

func TestExpand(t *testing.T) {
	var tests = []struct {
		input string
		want  string
	}{
		{"", ""},
		{"test test $test test", "test test TEST test"},
		{"test", "test"},
		{"$test", "TEST"},
	}

	for _, test := range tests {
		got := expand(test.input, toUpper)
		if got != test.want {
			t.Errorf("want is \"%s\", but got is \"%s\"\n", test.want, got)
		}
	}
}
