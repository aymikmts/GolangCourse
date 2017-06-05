package intset

import "testing"

// データセットを作成する(ex01_test.goに定義)
//func makeDataSet(s []int) *IntSet {
//	var set IntSet
//	for _, v := range s {
//		set.Add(v)
//	}
//	return &set
//}

func TestAddAll(t *testing.T) {
	var tests = []struct {
		input []int
		len   int
		want  []int
	}{
		{[]int{}, 0, []int{}},
		{[]int{1}, 1, []int{1}},
		{[]int{1, 10, 100}, 3, []int{1, 10, 100}},
	}
	for _, test := range tests {
		want := makeDataSet(test.want)

		var s IntSet
		s.AddAll(test.input...)
		if s.Len() != test.len {
			t.Errorf("got length is %d, but want is %d.\n", s.Len(), test.len)
		}
		if s.String() != want.String() {
			t.Errorf("got is %s, but want is %s.\n", s.String(), want.String())
		}
	}
}
