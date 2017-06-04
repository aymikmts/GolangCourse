package main

import "testing"

func TestMaxWithErr(t *testing.T) {
	var tests = []struct {
		in    []int
		out   int
		isErr bool
	}{
		{[]int{}, 0, true},
		{[]int{1}, 1, false},
		{[]int{1, 2, 3, 4, 5}, 5, false},
		{[]int{2, 3, 4, 5, 1}, 5, false},
	}

	for _, test := range tests {
		got, err := maxWithErr(test.in...)
		if test.isErr {
			if err == nil {
				t.Errorf("err is expected, but got is nil\n")
			}
			continue
		} else {
			if err != nil {
				t.Errorf("err is not expected, but got is %v\n", err)
				continue
			}
		}

		if got != test.out {
			t.Errorf("got is %v, but want is %v.\n", got, test.out)
		}
	}
}

func TestMax(t *testing.T) {
	var tests = []struct {
		in  []int
		out int
	}{
		{[]int{1}, 1},
		{[]int{1, 2, 3, 4, 5}, 5},
		{[]int{2, 3, 4, 5, 1}, 5},
	}

	for _, test := range tests {
		got := max(test.in[0], test.in[1:]...)
		if got != test.out {
			t.Errorf("got is %v, but want is %v.\n", got, test.out)
		}
	}
}
