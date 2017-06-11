// Ex08は入力キーに対して多段ソートを行います。
package main

import (
	"fmt"
	"os"
	"sort"

	"../sorting"
)

var tracks = []*sorting.Track{
	{"Go", "Delilah", "From the Roots Up", 2012, sorting.Length("3m38s")},
	{"Go", "Moby", "Moby", 1992, sorting.Length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, sorting.Length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, sorting.Length("4m24s")},
}

func main() {
	keys := os.Args[1:]
	fmt.Printf("keys: %v\n", keys)
	sort.Sort(sorting.MultiSort(tracks, keys))

	sorting.PrintTracks(tracks)
}
