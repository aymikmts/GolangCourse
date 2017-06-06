package treesort

import "testing"

func TestString(t *testing.T) {
	var tests = []struct {
		input []int
		want  string
	}{
		{[]int{1}, "{1}"},
		{[]int{1, 1, 1, 1, 1}, "{1 1 1 1 1}"},
		{[]int{1, 2, 3, 4, 5}, "{1 2 3 4 5}"},
		{[]int{5, 4, 3, 2, 1}, "{1 2 3 4 5}"},
		{[]int{3, 4, 2, 1, 5}, "{1 2 3 4 5}"},
	}

	for _, test := range tests {
		var ts *tree
		for _, v := range test.input {
			ts = add(ts, v)
		}
		if ts.String() != test.want {
			t.Errorf("got is %v, but want is %v\n", ts.String(), test.want)
		}
	}
}
