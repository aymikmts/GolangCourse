package echo2

import "testing"

func BenchmarkEcho(b *testing.B) {
	args := []string{"test1", "test2", "test3"}
	for i := 0; i < b.N; i++ {
		Echo(args[:])
	}
}
