package intset

import (
	"testing"
)

// データセットを作成する(ex01_test.goに定義)
//func makeDataSet(s []int) *IntSet {
//	var set IntSet
//	for _, v := range s {
//		set.Add(v)
//	}
//	return &set
//}

func TestElems(t *testing.T) {
	var tests = []struct {
		input []int
	}{
		{[]int{}},
		{[]int{1}},
		{[]int{0, 1, 10, 100}},
		{[]int{1, 1 << 1, 1 << 2, 1 << 3}},
	}

	for _, test := range tests {
		input := makeDataSet(test.input)
		got := input.Elems()
		if len(got) != len(test.input) {
			t.Errorf("got size is %d, but want is %d.\n", len(got), len(test.input))
		}
		for _, v := range got {
			if !input.Has(v) {
				t.Errorf("input.Has(%v) is false, but want is true.\n", v)
			}
		}
	}
}
