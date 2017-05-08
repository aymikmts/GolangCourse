package rev

import (
	"testing"
)

func TestReverseByPointer(t *testing.T) {
	var tests = []struct {
		array [6]int
		want  [6]int
	}{
		{[6]int{0, 1, 2, 3, 4, 5}, [6]int{5, 4, 3, 2, 1, 0}},
		{[6]int{0, 1, 2}, [6]int{0, 0, 0, 2, 1, 0}},
		{[...]int{0, 1, 2, 3, 4, 5}, [6]int{5, 4, 3, 2, 1, 0}},
	}
	for _, test := range tests {
		input := test.array
		ReverseByPointer(&test.array)
		if test.array != test.want {
			t.Errorf("ReverseByPointer(%v) : %v", input, test.array)
		}
	}
}
