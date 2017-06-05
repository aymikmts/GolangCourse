package intset

import "testing"

func TestAddAll(t *testing.T) {
	var tests = []struct {
		input []int
		len   int
		want  string
	}{
		{[]int{}, 0, "{}"},
		{[]int{1}, 1, "{1}"},
		{[]int{1, 10, 100}, 3, "{1 10 100}"},
	}
	for _, test := range tests {
		var s IntSet
		s.AddAll(test.input...)
		if s.Len() != test.len {
			t.Errorf("got length is %d, but want is %d.\n", s.Len(), test.len)
		}
		if s.String() != test.want {
			t.Errorf("got is %s, but want is %s.\n", s.String(), test.want)
		}
	}
}
