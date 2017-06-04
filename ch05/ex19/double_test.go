package main

import "testing"

func TestDouble(t *testing.T) {
	var tests = []struct {
		input int
		want  int
	}{
		{0, 0},
		{1, 2},
		{1000, 2000},
	}
	for _, test := range tests {
		got := double(test.input)
		if got != test.want {
			t.Errorf("got is %d, but want is %d\n", got, test.want)
		}
	}
}
