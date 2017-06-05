package intset

import "testing"

// データセットを作成する
func makeDataSet(s []int) *IntSet {
	var set IntSet
	for _, v := range s {
		set.Add(v)
	}
	return &set
}

func TestLen(t *testing.T) {
	var tests = []struct {
		input []int
		want  int
	}{
		{[]int{}, 0},
		{[]int{1, 10, 100}, 3},
		{[]int{0}, 1},
	}
	for _, test := range tests {
		input := makeDataSet(test.input)
		got := input.Len()
		if got != test.want {
			t.Errorf("got is %d, want is %d.\n", got, test.want)
		}
	}
}

func TestRemove(t *testing.T) {
	var tests = []struct {
		input  []int
		remove int
		want   []int
	}{
		{[]int{1, 10, 100}, 1, []int{10, 100}},
		{[]int{1, 10, 100}, 2, []int{1, 10, 100}},
	}
	for _, test := range tests {
		input := makeDataSet(test.input)
		want := makeDataSet(test.want)
		input.Remove(test.remove)
		if input.String() != want.String() {
			t.Errorf("got is %s, want is %s.\n", input.String(), want.String())
		}
	}
}

func TestClear(t *testing.T) {
	var tests = []struct {
		input []int
		want  []int
	}{
		{[]int{}, []int{}},
		{[]int{1}, []int{}},
		{[]int{1, 10, 100}, []int{}},
	}
	for _, test := range tests {
		input := makeDataSet(test.input)
		want := makeDataSet(test.want)
		input.Clear()
		if input.String() != want.String() {
			t.Errorf("got is %s, want is %s.\n", input.String(), want.String())
		}
		if want.Len() != 0 {
			t.Errorf("got is %d, want is 0.\n", want.Len())
		}
	}
}

func TestCopy(t *testing.T) {
	var tests = []struct {
		input []int
		want  []int
	}{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{1, 10, 100}, []int{1, 10, 100}},
	}
	for _, test := range tests {
		input := makeDataSet(test.input)
		want := makeDataSet(test.want)
		got := input.Copy()
		if got.String() != want.String() {
			t.Errorf("got is %s, want is %s.\n", got.String(), want.String())
		}
		if got.Len() != want.Len() {
			t.Errorf("got is %d, want is %d.\n", got.Len(), want.Len())
		}
	}
}
