package dedup

import (
	"reflect"
	"testing"
)

func TestDedupSpace(t *testing.T) {
	var tests = []struct {
		input []byte
		want  []byte
	}{
		{[]byte("こんにちは 世界"), []byte("こんにちは 世界")},    // 重複なしケース(ASCIIスペース)
		{[]byte("こんにちは　世界"), []byte("こんにちは　世界")},    // 重複なしケース(全角スペース)
		{[]byte("こんにちは　　世界"), []byte("こんにちは 世界")},   // 全角スペース重複
		{[]byte("こんにちは\n\t世界"), []byte("こんにちは 世界")}, // unicode.IsSpace対応文字
		{[]byte("　こんにちは世界"), []byte("　こんにちは世界")},    // 文字列の始めに空白文字
		{[]byte("　　こんにちは世界"), []byte(" こんにちは世界")},
		{[]byte("\t\tこんにちは世界"), []byte(" こんにちは世界")},
		{[]byte("こんにちは世界　"), []byte("こんにちは世界　")}, // 文字列の終わりに空白文字
		{[]byte("こんにちは世界　　"), []byte("こんにちは世界 ")},
		{[]byte("こんにちは世界\t\t"), []byte("こんにちは世界 ")},
	}
	for _, test := range tests {
		got := DedupSpace(test.input)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("DedupSpace(%q) = %q", test.input, got)
		}
	}
}
