package echo2

import (
	"strconv"
	"testing"
)

func BenchmarkEcho(b *testing.B) {
	var args []string
	for i := 0; i < 30; i++ {
		s := strconv.Itoa(i)
		args = append(args, s)
	}
	for i := 0; i < b.N; i++ {
		Echo(args[:])
	}
}
