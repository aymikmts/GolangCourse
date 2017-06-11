package sorting

import (
	"sort"
	"testing"
)

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, Length("3m38s")},
	{"Go", "Moby", "Moby", 1992, Length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, Length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, Length("4m24s")},
}

func TestMultiSort(t *testing.T) {
	var tests = []struct {
		keys []string
	}{
		{[]string{"Title"}},
		{[]string{"Title", "Year"}},
	}

	for _, test := range tests {
		sort.Sort(MultiSort(tracks, test.keys))
		PrintTracks(tracks)
	}

}
