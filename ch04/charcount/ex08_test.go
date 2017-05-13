package charcount

import (
	"testing"
)

func TestCharType(t *testing.T) {
	var tests = []struct {
		input rune
		want  string
	}{
		{'a', "letter"},
		{'あ', "letter"},
		{' ', "space"},
		{'1', "number"},
		{'\n', "control"},
		{'〠', "symbol"},
	}
	for _, test := range tests {
		got := charType(test.input)
		if got != test.want {
			t.Errorf(`charType(%v) = %v`, test.input, got)
		}
	}
}
