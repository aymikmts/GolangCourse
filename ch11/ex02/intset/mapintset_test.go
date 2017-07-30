package intset

import (
	"fmt"
	"testing"
)

func Example_MapIntSet_one() {
	var x, y MapIntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String())

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String())

	x.UnionWith(&y)
	fmt.Println(x.String())

	fmt.Println(x.Has(9), x.Has(123))

	// Output:
	// {1 9 144}
	// {9 42}
	// {1 9 42 144}
	// true false
}

func Example_MapIntSet_two() {
	var x MapIntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(42)

	fmt.Println(&x)         // "{1 9 42 144}"
	fmt.Println(x.String()) // "{1 9 42 144}"
	//fmt.Println(x)          // "{map[1:true 144:true 9:true 42:true]}"

	// Output:
	// {1 9 42 144}
	// {1 9 42 144}
}

//
// each function test
//
func TestHas_MapIntSet(t *testing.T) {
	// nil case
	var test MapIntSet
	if test.Has(1) {
		t.Errorf("testing nil case. got is true, but want is false.")
	}

	// exist case
	test.set = map[int]bool{1: true, 16: true, 32: true}
	if !test.Has(1) {
		t.Errorf("test.Has(1): got is false, but want is true.")
	}

	// no exist case
	if test.Has(2) {
		t.Errorf("test.Has(2): got is true, but want is false.")
	}
}

func TestAdd_MapIntSet(t *testing.T) {
	defer func() {
		if p := recover(); p != nil {
			t.Errorf("The code must not panic.")
		}
	}()

	// nil case
	var test MapIntSet
	test.Add(1)

	// exist case; Add 1
	test.set = make(map[int]bool)
	if test.set[1] {
		t.Errorf("test-map should not have 1.")
	}
	test.Add(1)
	if !test.set[1] {
		t.Errorf("test-map should have 1.")
	}

	// Add minus value
	test.Add(-1)
	if test.set[-1] {
		t.Errorf("test-map should not have -1.")
	}
}

func TestUnionWith_MapIntSet(t *testing.T) {
	defer func() {
		if p := recover(); p != nil {
			t.Errorf("The code must not panic.")
		}
	}()

	var x, y MapIntSet

	// nil case
	x.UnionWith(&y)

	// nil case
	x.set = map[int]bool{1: true, 9: true, 144: true}
	x.UnionWith(&y)

	//
	y.set = map[int]bool{9: true, 42: true}
	x.UnionWith(&y)

	keys := []int{1, 9, 42, 144}
	for _, k := range keys {
		if !x.set[k] {
			t.Errorf("x should have %d.", k)
		}
	}
}
