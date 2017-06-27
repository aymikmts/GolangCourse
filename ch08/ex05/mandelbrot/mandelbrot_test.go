package mandelbrot

import (
	"bytes"
	"testing"
)

func BenchmarkDrawWithoutGoroutine(b *testing.B) {
	out := new(bytes.Buffer)
	DrawWithoutGoroutine(out)
}

func BenchmarkDrawWithGoroutine(b *testing.B) {
	out := new(bytes.Buffer)
	DrawWithGoroutine(out)
}
