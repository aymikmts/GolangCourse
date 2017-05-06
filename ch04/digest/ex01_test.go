package digest

import (
	"crypto/sha256"
	"testing"
)

// 全てのbitが1の配列を用意する
var data1, data2 [32]uint8

func init() {
	for i := 0; i < 32; i++ {
		data1[i] = 0xFF
		data2[i] = 0xFF
	}
}

func TestCountDiffBit(t *testing.T) {
	var tests = []struct {
		x1   [32]uint8
		x2   [32]uint8
		want int
	}{
		{[32]uint8{}, [32]uint8{}, 0},
		{[32]uint8{31: 1}, [32]uint8{}, 1},
		{[32]uint8{31: 255}, [32]uint8{}, 8},
		{data1, data2, 0},
		{data1, [32]uint8{}, 256},
		{sha256.Sum256([]byte("x")), sha256.Sum256([]byte("x")), 0},
		{sha256.Sum256([]byte("x")), sha256.Sum256([]byte("X")), 125},
	}
	for _, test := range tests {
		got := CountDiffBit(test.x1, test.x2)
		if got != test.want {
			t.Errorf("CountDiffBit(%v, %v) = %v", test.x1, test.x2, got)
		}
	}
}
