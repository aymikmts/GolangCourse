package popcountroop

import "testing"

func TestPopCount(t *testing.T) {
	var tests = []struct {
		input uint64
		want  int
	}{
		{uint64(0), 0},   // 0b
		{uint64(8), 1},   // 1000b
		{uint64(10), 2},  // 1010b
		{uint64(15), 4},  // 1111b
		{uint64(255), 8}, // 11111111b
	}
	for _, test := range tests {
		if got := PopCount(test.input); got != test.want {
			t.Errorf(`PopCount(%v) = %v`, test.input, got)
		}
	}
}

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(uint64(i))
	}
}
