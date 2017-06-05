package main

import "testing"

// データセットを作成する
func makeDataSet(s []int) *IntSet {
	var set IntSet
	for _, v := range s {
		set.Add(v)
	}
	return &set
}

func TestHas(t *testing.T) {
	var tests = []struct {
		input []int
		val   int
		want  bool
	}{
		{[]int{}, 1, false},
		{[]int{1}, 1, true},
		{[]int{1}, 2, false},
		{[]int{1, 10, 100}, 10, true},
		{[]int{1, 10, 100}, 1000, false},
	}
	for _, test := range tests {
		input := makeDataSet(test.input)
		got := input.Has(test.val)
		if got != test.want {
			t.Errorf("input.Has(%v) got is %v, but want is %v.\n", test.input, got, test.want)
		}
	}
}

func TestAdd(t *testing.T) {
	var tests = []struct {
		val  []int
		len  int
		want string
	}{
		{[]int{}, 0, "{}"},
		{[]int{1}, 1, "{1}"},
		{[]int{1, 100}, 2, "{1 100}"},
	}
	for _, test := range tests {
		var s IntSet
		for _, v := range test.val {
			s.Add(v)
		}
		if s.Len() != test.len {
			t.Errorf("s.Len() got is %v, but want is %v.\n", s.Len(), test.len)
		}
		if s.String() != test.want {
			t.Errorf("got is %v, but want is %v.\n", s.String(), test.want)
		}
	}
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
