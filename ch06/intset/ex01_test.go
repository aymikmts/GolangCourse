package intset

import "testing"

func TestLen(t *testing.T) {
	var data1, data2, data3 IntSet
	data2.Add(1)
	data2.Add(10)
	data2.Add(100)
	data3.Add(0)

	var tests = []struct {
		input IntSet
		want  int
	}{
		{data1, 0},
		{data2, 3},
		{data3, 1},
	}
	for _, test := range tests {
		got := test.input.Len()
		if got != test.want {
			t.Errorf("got is %d, want is %d\n", got, test.want)
		}
	}
}

func TestRemove(t *testing.T) {
	var data1, data2 IntSet
	data1.Add(1)
	data1.Add(10)
	data1.Add(100)
	data2 = *(data1.Copy())

	var want1, want2 IntSet
	want1.Add(10)
	want1.Add(100)
	want2.Add(1)
	want2.Add(10)
	want2.Add(100)

	var tests = []struct {
		input  IntSet
		remove int
		want   IntSet
	}{
		{data1, 1, want1},
		{data2, 2, want2},
	}
	for _, test := range tests {
		test.input.Remove(test.remove)
		if test.input.String() != test.want.String() {
			t.Errorf("got is %s, want is %s\n", test.input.String(), test.want.String())
		}
	}
}

func TestClear(t *testing.T) {
	var data, want, empty IntSet
	data.Add(1)
	data.Add(10)
	data.Add(100)

	var tests = []struct {
		input IntSet
		want  IntSet
	}{
		{data, want},
		{empty, want},
	}
	for _, test := range tests {
		test.input.Clear()
		if test.input.String() != test.want.String() {
			t.Errorf("got is %s, want is %s\n", test.input.String(), test.want.String())
		}
		if test.want.Len() != 0 {
			t.Errorf("got is 0, want is %d\n", test.want.Len())
		}
	}
}

func TestCopy(t *testing.T) {
	var data, want, empty IntSet
	data.Add(1)
	data.Add(10)
	data.Add(100)

	want.Add(1)
	want.Add(10)
	want.Add(100)

	var tests = []struct {
		input IntSet
		want  IntSet
	}{
		{data, want},
		{empty, empty},
	}
	for _, test := range tests {
		got := *(test.input.Copy())
		if test.want.String() != got.String() {
			t.Errorf("got is %s, want is %s\n", got.String(), test.want.String())
		}
		if test.want.Len() != got.Len() {
			t.Errorf("got is %d, want is %d\n", got.Len(), test.want.Len())
		}
	}
}
