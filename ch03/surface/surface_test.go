package surface

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

func TestPrintXML(t *testing.T) {
	out := new(bytes.Buffer)
	PrintXML(out)

	n, err := out.WriteTo(os.Stdout)

	if err != nil {
		fmt.Fprintf(os.Stderr, "err:%d\n", err)
	}
	fmt.Printf("\nn=%d\n", n)
}

func TestCorner(t *testing.T) {
	var tests = []struct {
		i  int
		j  int
		ok bool
	}{
		{0, 0, true},
		{50, 50, false},
	}
	for _, test := range tests {
		if _, _, got := corner(test.i, test.j); got != test.ok {
			t.Errorf(`corner(%v, %v) = %v`, test.i, test.j, got)
		}
	}
}

// func BenchmarkPopCount(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		PopCount(uint64(i))
// 	}
// }
