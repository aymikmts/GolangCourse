package dedup

import (
	"reflect"
	"testing"
)

func TestDedup(t *testing.T) {
	var tests = []struct {
		input []string
		want  []string
	}{
		{[]string{"aa", "bb", "cc", "dd"}, []string{"aa", "bb", "cc", "dd"}}, // 重複なしケース
		{[]string{"aa", "AA", "bb", "cc"}, []string{"aa", "AA", "bb", "cc"}}, // 大文字・小文字を区別する
		{[]string{"aa", "aa", "aa", "bb"}, []string{"aa", "bb"}},             // 最初3要素が重複
		{[]string{"aa", "bb", "bb", "bb"}, []string{"aa", "bb"}},             // 後ろ3要素が重複
		{[]string{"aa", "aa", "aa", "aa"}, []string{"aa"}},                   // 全ての要素が重複
	}
	for _, test := range tests {
		got := Dedup(test.input)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("Dedup(%v) = %v", test.input, got)
		}
	}
}
