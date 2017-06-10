package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestCountingWriter(t *testing.T) {
	var tests = []struct {
		input []string
	}{
		{[]string{"", "", ""}},
		{[]string{"a", "b", "c"}},
		{[]string{"test", "Hello, World!", "こんにちは"}},
	}
	for _, test := range tests {
		buf := new(bytes.Buffer)
		w, cnt := CountingWriter(buf)

		var got int
		for _, s := range test.input {
			got += len(s)
			fmt.Fprintf(w, "%s", s)
		}

		if got != int(*cnt) {
			t.Errorf("got is %d, but want is %d.\n", got, *cnt)
		}
	}
}
