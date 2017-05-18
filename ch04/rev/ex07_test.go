package rev

import (
	"bytes"
	"testing"
)

func TestReverseUnicodeByte(t *testing.T) {
	var tests = []struct {
		input []byte
		want  []byte
	}{
		{[]byte("test"), []byte("tset")},   // 1byte
		{[]byte("ĀĐĠİ"), []byte("İĠĐĀ")},   // 2byte
		{[]byte("こんにちは"), []byte("はちにんこ")}, // 3byte
		{[]byte("你再割卿"), []byte("卿割再你")},   // 4byte
		{[]byte("tĀこ你"), []byte("你こĀt")},   // ごちゃまぜ
	}
	for _, test := range tests {
		input := test.input
		//fmt.Printf("input:\n%08b\n", input)
		ReverseUnicodeByte(test.input)
		//fmt.Printf("output:\n%08b\n\n", test.input)
		if !bytes.Equal(test.input, test.want) {
			t.Errorf("ReverseUnicodeByte(%v) : %v", input, test.input)
		}
	}
}
