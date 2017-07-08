package popcount

import (
	"fmt"
	"sync"
	"testing"
)

func TestInitTable(t *testing.T) {
	PopCount(0x123456789)
}

func TestPopCount(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(i int) {
			ret := PopCount(0x123456789)
			fmt.Printf("[%2d]PopCount(%08x)=%d\n", i, i, ret)
			wg.Done()
		}(i)
	}
	wg.Wait()

	fmt.Printf("done.\n")
}
