package mandelbrot

import (
	"testing"
)

func BenchmarkFormatCmplx64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		formatCmplx64(100, 100)
	}
}

func BenchmarkFormatCmplx128(b *testing.B) {
	for i := 0; i < b.N; i++ {
		formatCmplx128(100, 100)
	}
}

func BenchmarkFormatBigFloat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		formatBigFloat(100, 100)
	}
}

func BenchmarkFormatBigRat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		formatBigRat(100, 100)
	}
}
