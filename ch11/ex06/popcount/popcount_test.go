package popcount

import "testing"

func benchmarkPopCount(b *testing.B, val uint64) {
	for i := 0; i < b.N; i++ {
		PopCount(val)
	}
}

func benchmarkPopCountByShifting(b *testing.B, val uint64) {
	for i := 0; i < b.N; i++ {
		PopCountByShifting(val)
	}
}

func benchmarkPopCountByClearing(b *testing.B, val uint64) {
	for i := 0; i < b.N; i++ {
		PopCountByClearing(val)
	}
}

// input value: 0
func BenchmarkPopCount_0(b *testing.B)            { benchmarkPopCount(b, 0) }
func BenchmarkPopCountByShifiting_0(b *testing.B) { benchmarkPopCountByShifting(b, 0) }
func BenchmarkPopCountByClearing_0(b *testing.B)  { benchmarkPopCountByClearing(b, 0) }

// input value: 0xFFFF(16 count)
func BenchmarkPopCount_16c(b *testing.B)            { benchmarkPopCount(b, 0xFFFF) }
func BenchmarkPopCountByShifiting_16c(b *testing.B) { benchmarkPopCountByShifting(b, 0xFFFF) }
func BenchmarkPopCountByClearing_16c(b *testing.B)  { benchmarkPopCountByClearing(b, 0xFFFF) }

// input value: 0xFFFFFFFF(32 count)
func BenchmarkPopCount_32c(b *testing.B)            { benchmarkPopCount(b, 0xFFFFFFFF) }
func BenchmarkPopCountByShifiting_32c(b *testing.B) { benchmarkPopCountByShifting(b, 0xFFFFFFFF) }
func BenchmarkPopCountByClearing_32c(b *testing.B)  { benchmarkPopCountByClearing(b, 0xFFFFFFFF) }

// input value: 0xFFFFFFFFFFFFFFFF(64 count)
func BenchmarkPopCount_64c(b *testing.B) { benchmarkPopCount(b, 0xFFFFFFFFFFFFFFFF) }
func BenchmarkPopCountByShifiting_64c(b *testing.B) {
	benchmarkPopCountByShifting(b, 0xFFFFFFFFFFFFFFFF)
}
func BenchmarkPopCountByClearing_64c(b *testing.B) {
	benchmarkPopCountByClearing(b, 0xFFFFFFFFFFFFFFFF)
}

// Go 1.8, 1.3GHz Intel Core i5
// BenchmarkPopCount_0-4                   2000000000               0.41 ns/op
// BenchmarkPopCountByShifiting_0-4        20000000                98.9 ns/op
// BenchmarkPopCountByClearing_0-4         1000000000               2.49 ns/op
//
// BenchmarkPopCount_16c-4                 2000000000               0.45 ns/op
// BenchmarkPopCountByShifiting_16c-4      20000000               101 ns/op
// BenchmarkPopCountByClearing_16c-4       100000000               15.0 ns/op
//
// BenchmarkPopCount_32c-4                 2000000000               0.42 ns/op
// BenchmarkPopCountByShifiting_32c-4      20000000               106 ns/op
// BenchmarkPopCountByClearing_32c-4       50000000                29.4 ns/op
//
// BenchmarkPopCount_64c-4                 2000000000               0.41 ns/op
// BenchmarkPopCountByShifiting_64c-4      20000000                96.1 ns/op
// BenchmarkPopCountByClearing_64c-4       20000000                69.3 ns/op
