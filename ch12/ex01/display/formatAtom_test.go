package display

import (
	"reflect"
	"testing"
)

type A struct {
	id   int
	name string
}

type B struct {
	id      int
	name    string
	structA A
}

func TestFormatAtom(t *testing.T) {
	tests := []struct {
		val interface{}
		out string
	}{
		{int(10), "10"},
		{uint(10), "10"},
		{true, "true"},
		{"string", "\"string\""},

		// struct case
		{A{id: 0, name: "a"}, "A{id: 0, name: \"a\"}"},
		// inner struct case
		{B{id: 0, name: "b", structA: A{id: 0, name: "a"}},
			"B{id: 0, name: \"b\", structA: A{id: 0, name: \"a\"}}"},

		// array case
		{[3]int{10, 20, 30}, "[3]int{10, 20, 30}"},
		{[3]string{"a", "b", "c"}, "[3]string{\"a\", \"b\", \"c\"}"},
	}

	for _, test := range tests {
		got := formatAtom(reflect.ValueOf(test.val))
		if got != test.out {
			t.Errorf("formatAtom(%v) = %s, want %s", test.val, got, test.out)
		}
	}
}
