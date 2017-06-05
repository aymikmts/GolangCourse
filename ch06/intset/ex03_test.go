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

func TestIntersectWith(t *testing.T) {
	var tests = []struct {
		x    []int
		y    []int
		want []int
	}{
		{[]int{}, []int{}, []int{}},
		{[]int{1, 65, 129}, []int{1, 65, 129}, []int{1, 65, 129}}, // すべて一致
		{[]int{1, 65, 129}, []int{2, 66, 130}, []int{}},           // すべて不一致
		{[]int{1, 65, 129}, []int{2, 65, 130}, []int{65}},         // 一つ一致
		{[]int{1, 65, 129}, []int{2, 65, 129}, []int{65, 129}},    // 二つ一致
		{[]int{1, 65, 129}, []int{2, 65, 130, 193}, []int{65}},    // tの方がサイズ大きい
		{[]int{1, 65, 129, 193}, []int{2, 65, 130}, []int{65}},    // sの方がサイズ大きい
	}

	for _, test := range tests {
		x := makeDataSet(test.x)
		y := makeDataSet(test.y)
		want := makeDataSet(test.want)

		x.IntersectWith(y)
		if x.String() != want.String() {
			t.Errorf("got is %v, but want is %v.\n", x.String(), want.String())
		}
	}
}

func TestDifferenceWith(t *testing.T) {
	var tests = []struct {
		x    []int
		y    []int
		want []int
	}{
		{[]int{}, []int{}, []int{}},
		{[]int{1, 65, 129}, []int{1, 65, 129}, []int{}},                 // すべて一致
		{[]int{1, 65, 129}, []int{2, 66, 130}, []int{1, 65, 129}},       // すべて不一致
		{[]int{1, 65, 129}, []int{2, 65, 130}, []int{1, 129}},           // 一つ一致
		{[]int{1, 65, 129}, []int{2, 65, 129}, []int{1}},                // 二つ一致
		{[]int{1, 65, 129}, []int{2, 65, 130, 193}, []int{1, 129}},      // tの方がサイズ大きい
		{[]int{1, 65, 129, 193}, []int{2, 65, 130}, []int{1, 129, 193}}, // sの方がサイズ大きい
	}

	for _, test := range tests {
		x := makeDataSet(test.x)
		y := makeDataSet(test.y)
		want := makeDataSet(test.want)

		x.DifferenceWith(y)
		if x.String() != want.String() {
			t.Errorf("got is %v, but want is %v.\n", x.String(), want.String())
		}
	}
}

func TestSymmetricDifference(t *testing.T) {
	var tests = []struct {
		x    []int
		y    []int
		want []int
	}{
		{[]int{}, []int{}, []int{}},
		{[]int{1, 65, 129}, []int{1, 65, 129}, []int{}},                         // すべて一致
		{[]int{1, 65, 129}, []int{2, 66, 130}, []int{1, 2, 65, 66, 129, 130}},   // すべて不一致
		{[]int{1, 65, 129}, []int{2, 65, 130}, []int{1, 2, 129, 130}},           // 一つ一致
		{[]int{1, 65, 129}, []int{2, 65, 129}, []int{1, 2}},                     // 二つ一致
		{[]int{1, 65, 129}, []int{2, 65, 130, 193}, []int{1, 2, 129, 130, 193}}, // tの方がサイズ大きい
		{[]int{1, 65, 129, 193}, []int{2, 65, 130}, []int{1, 2, 129, 130, 193}}, // sの方がサイズ大きい
	}

	for _, test := range tests {
		x := makeDataSet(test.x)
		y := makeDataSet(test.y)
		want := makeDataSet(test.want)

		x.SymmetricDifference(y)
		if x.String() != want.String() {
			t.Errorf("got is %v, but want is %v.\n", x.String(), want.String())
		}
	}
}
