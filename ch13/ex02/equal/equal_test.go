package equal

import (
	"fmt"
	"testing"
)

func TestEqualNanoScale(t *testing.T) {
	tests := []struct {
		x, y interface{}
		want bool
	}{
		{float64(1.0), float64(1.0 + 0.9e-9), true},
		{float64(1.0), float64(1.0 + 1.0e-9), false},
		{float64(1.0), float64(1.0 + 1.1e-9), false},
		{complex128(1.0 + 1.0i), complex128((1.0 + 0.9e-9) + 1.0i), true},
		{complex128(1.0 + 1.0i), complex128((1.0 + 1.0e-9) + 1.0i), false},
		{complex128(1.0 + 1.0i), complex128((1.0 + 1.1e-9) + 1.0i), false},
	}
	for _, test := range tests {
		got := Equal(test.x, test.y)
		if got != test.want {
			t.Errorf("Equal(%#v, %#v) got %v, want %v", test.x, test.y, got, test.want)
		}
	}
}
func Example_equal() {
	fmt.Println(Equal([]int{1, 2, 3}, []int{1, 2, 3}))
	fmt.Println(Equal([]string{"foo"}, []string{"bar"}))
	fmt.Println(Equal([]string(nil), []string{}))
	fmt.Println(Equal(map[string]int(nil), map[string]int{}))

	// Output:
	// true
	// false
	// true
	// true
}

func Example_equalCycle() {
	// 循環したリンクリスト a -> b -> a と c -> c
	type link struct {
		value string
		tail  *link
	}
	a, b, c := &link{value: "a"}, &link{value: "b"}, &link{value: "c"}
	a.tail, b.tail, c.tail = b, a, c
	fmt.Println(Equal(a, a))
	fmt.Println(Equal(b, b))
	fmt.Println(Equal(c, c))
	fmt.Println(Equal(a, b))
	fmt.Println(Equal(a, c))

	// Output:
	// true
	// true
	// true
	// false
	// false
}
