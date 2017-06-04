package main

import "testing"

func TestJoinByVariable(t *testing.T) {
	var tests = []struct {
		input []string
		want  string
		isErr bool
	}{
		{[]string{}, "", true},
		{[]string{","}, "", true},
		{[]string{"a", "b", "c", ","}, "a,b,c", false},
		{[]string{"a", "b", "c", "d", "e", "/"}, "a/b/c/d/e", false},
	}

	for _, test := range tests {
		got, err := joinByVariable(test.input...)
		if test.isErr {
			if err == nil {
				t.Errorf("error is expected, but got is nil.\n")
			}
			continue
		}
		if got != test.want {
			t.Errorf("got is \"%s\", but want is \"%s\"\n", got, test.want)
		}
	}
}
