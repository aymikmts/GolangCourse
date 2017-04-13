package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestEcho(t *testing.T) {
	var tests = []struct {
		args []string
		want string
	}{
		{[]string{"path/main.exe", "test1", "test2", "test3"}, "main.exe test1 test2 test3\n"},
		{[]string{""}, "\n"},
	}

	for _, test := range tests {
		descr := fmt.Sprintf("echo(%v)", test.args)

		out = new(bytes.Buffer)
		if err := echo(test.args); err != nil {
			t.Errorf("%s failed; %v", descr, err)
			continue
		}
		got := out.(*bytes.Buffer).String()
		ret := strings.Contains(got, test.want)
		if !ret {
			t.Errorf("%s = %q, want %q", descr, got, test.want)
		}
	}
}
