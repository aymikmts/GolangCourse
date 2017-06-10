package main

import (
	"fmt"
	"testing"
)

func TestWordCounterWrite(t *testing.T) {
	var tests = []struct {
		input []byte
		want  int
	}{
		{[]byte(""), 0},
		{[]byte("test"), 1},
		{[]byte("test1 test2"), 2},
		{[]byte("test1\ntest2"), 2},
		{[]byte("test1 test2\ntest3 test4\n"), 4},
	}

	for _, test := range tests {
		var c WordCounter
		c.Write(test.input)
		if int(c) != test.want {
			t.Errorf("input: %v\ngot is %d, but want is %d\n.", test.input, c, test.want)
		}

		c = 0
		fmt.Fprintf(&c, "%s", test.input)
		if int(c) != test.want {
			t.Errorf("input: %v\ngot is %d, but want is %d\n.", test.input, c, test.want)
		}
	}
}

func TestLineCounterWrite(t *testing.T) {
	var tests = []struct {
		input []byte
		want  int
	}{
		{[]byte(""), 0},
		{[]byte("test"), 1},
		{[]byte("test1 test2"), 1},
		{[]byte("test1\ntest2"), 2},
		{[]byte("test1 test2\ntest3 test4\n"), 2},
	}

	for _, test := range tests {
		var c LineCounter
		c.Write(test.input)
		if int(c) != test.want {
			t.Errorf("input: %v\ngot is %d, but want is %d\n.", test.input, c, test.want)
		}

		c = 0
		fmt.Fprintf(&c, "%s", test.input)
		if int(c) != test.want {
			t.Errorf("input: %v\ngot is %d, but want is %d\n.", test.input, c, test.want)
		}
	}
}
