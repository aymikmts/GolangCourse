package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestEcho(t *testing.T) {
	var tests = []struct {
		args []string
		want string
	}{
		{[]string{"dummy.exe", "test1", "test2", "test3"}, "1 test1\n2 test2\n3 test3\n"},
		{[]string{""}, ""},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("echo(%v)", test.args)

		out = new(bytes.Buffer)
		if err := echo(test.args); err != nil {
			t.Errorf("%s failed; %v", descr, err)
			continue
		}
		got := out.(*bytes.Buffer).String()
		if got != test.want {
			t.Errorf("%s = %q, want %q", descr, got, test.want)
		}
	}
}
