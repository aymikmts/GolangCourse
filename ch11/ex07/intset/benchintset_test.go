package intset

import (
	"math"
	"math/rand"
	"testing"
	"time"
)

const addCnt = 100

// ワードの大きさの比較
// Add
func benchmark_IntSet_Add(b *testing.B, max int) {
	var x IntSet
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < b.N; i++ {
		for j := 0; j < addCnt; j++ {
			x.Add(rand.Intn(max))
		}
	}
}
func Benchmark_IntSet_Add_Word8(b *testing.B) {
	benchmark_IntSet_Add(b, math.MaxInt8)
}
func Benchmark_IntSet_Add_Word16(b *testing.B) {
	benchmark_IntSet_Add(b, math.MaxInt16)
}
func Benchmark_IntSet_Add_Word32(b *testing.B) {
	benchmark_IntSet_Add(b, math.MaxInt32)
}

// UnionWith
func benchmark_IntSet_UnionWith(b *testing.B, max int) {
	var x, y IntSet
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < b.N; i++ {
		for j := 0; j < addCnt; j++ {
			x.Add(rand.Intn(max))
			y.Add(rand.Intn(max))
		}
		x.UnionWith(&y)
	}
}
func Benchmark_IntSet_UnionWith_Word8(b *testing.B) {
	benchmark_IntSet_UnionWith(b, math.MaxInt8)
}
func Benchmark_IntSet_UnionWith_Word16(b *testing.B) {
	benchmark_IntSet_UnionWith(b, math.MaxInt16)
}
func Benchmark_IntSet_UnionWith_Word32(b *testing.B) {
	benchmark_IntSet_UnionWith(b, math.MaxInt32)
}

// IntersectWith
func benchmark_IntSet_IntersectWith(b *testing.B, max int) {
	var x, y IntSet
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < b.N; i++ {
		for j := 0; j < addCnt; j++ {
			x.Add(rand.Intn(max))
			y.Add(rand.Intn(max))
		}
		x.IntersectWith(&y)
	}
}
func Benchmark_IntSet_IntersectWith_Word8(b *testing.B) {
	benchmark_IntSet_IntersectWith(b, math.MaxInt8)
}
func Benchmark_IntSet_IntersectWith_Word16(b *testing.B) {
	benchmark_IntSet_IntersectWith(b, math.MaxInt16)
}
func Benchmark_IntSet_IntersectWith_Word32(b *testing.B) {
	benchmark_IntSet_IntersectWith(b, math.MaxInt32)
}

// 組み込みマップ型との比較
func Benchmark_IntSet_Add(b *testing.B) {
	var x IntSet
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < b.N; i++ {
		for j := 0; j < addCnt; j++ {
			x.Add(rand.Intn(math.MaxInt16))
		}
	}
}

func Benchmark_MapIntSet_Add(b *testing.B) {
	var x MapIntSet
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < b.N; i++ {
		for j := 0; j < addCnt; j++ {
			x.Add(rand.Intn(math.MaxInt16))
		}
	}
}

func Benchmark_IntSet_UnionWith(b *testing.B) {
	var x, y IntSet
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < b.N; i++ {
		for j := 0; j < addCnt; j++ {
			x.Add(rand.Intn(math.MaxInt16))
			y.Add(rand.Intn(math.MaxInt16))
		}
		x.UnionWith(&y)
	}
}

func Benchmark_MapIntSet_UnionWith(b *testing.B) {
	var x, y MapIntSet
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < b.N; i++ {
		for j := 0; j < addCnt; j++ {
			x.Add(rand.Intn(math.MaxInt16))
			y.Add(rand.Intn(math.MaxInt16))
		}
		x.UnionWith(&y)
	}
}

// Go1.8, 1.3GHz Intel Core i5
// Benchmark_IntSet_Add_Word8-4                      300000              5527 ns/op
// Benchmark_IntSet_Add_Word16-4                     200000              5883 ns/op
// Benchmark_IntSet_Add_Word32-4                      50000             25723 ns/op
//
// Benchmark_IntSet_UnionWith_Word8-4                100000             11356 ns/op
// Benchmark_IntSet_UnionWith_Word16-4               100000             12229 ns/op
// Benchmark_IntSet_UnionWith_Word32-4                    2         500061504 ns/op
//
// Benchmark_IntSet_IntersectWith_Word8-4            100000             11622 ns/op
// Benchmark_IntSet_IntersectWith_Word16-4           100000             12290 ns/op
// Benchmark_IntSet_IntersectWith_Word32-4               10         130411238 ns/op
//
// Benchmark_IntSet_Add-4                            200000              5759 ns/op
// Benchmark_MapIntSet_Add-4                         100000             11849 ns/op
// Benchmark_IntSet_UnionWith-4                      100000             12105 ns/op
// Benchmark_MapIntSet_UnionWith-4                     3000           2350935 ns/op
