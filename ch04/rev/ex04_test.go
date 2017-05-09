package rev

import (
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {
	var tests = []struct {
		slice []int
		rot   int
		want  []int
	}{
		{[]int{10, 20, 30, 40, 50, 60}, 2, []int{30, 40, 50, 60, 10, 20}},  // 左へ2要素分回転
		{[]int{10, 20, 30, 40, 50, 60}, -2, []int{50, 60, 10, 20, 30, 40}}, // 右へ2要素分回転
		{[]int{10, 20, 30, 40, 50, 60}, 9, []int{40, 50, 60, 10, 20, 30}},  // スライス長以上の回転要素数を指定された場合のテスト
		{[]int{10, 20, 30, 40, 50, 60}, -9, []int{40, 50, 60, 10, 20, 30}},
		{[]int{10, 20, 30, 40, 50, 60}, 0, []int{10, 20, 30, 40, 50, 60}},  // 0回転
		{[]int{10, 20, 30, 40, 50, 60}, 12, []int{10, 20, 30, 40, 50, 60}}, // 0回転
		{[]int{10, 20, 30}, 2, []int{30, 10, 20}},                          // スライス要素数を変更してテスト
		{[]int{10, 20, 30}, -2, []int{20, 30, 10}},
		{[]int{10, 20, 30}, 9, []int{10, 20, 30}},
		{[]int{10, 20, 30}, -9, []int{10, 20, 30}},
	}
	for _, test := range tests {
		input := test.slice
		Rotate(test.slice, test.rot)
		if !reflect.DeepEqual(test.slice, test.want) {
			t.Errorf("Rotate(%v, %v) : %v", input, test.rot, test.slice)
		}
	}
}
