package sorting

import "sort"

type multiSort struct {
	t    []*Track
	keys []string
}

func (x multiSort) Len() int           { return len(x.t) }
func (x multiSort) Less(i, j int) bool { return less(x.t[i], x.t[j], x.keys) }
func (x multiSort) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }

func less(x, y *Track, keys []string) bool {
	for _, key := range keys {
		switch key {
		case "Title":
			if x.Title != y.Title {
				return x.Title < y.Title
			}
		case "Artist":
			if x.Artist != y.Artist {
				return x.Artist < y.Artist
			}
		case "Album":
			if x.Album != y.Album {
				return x.Album < y.Album
			}
		case "Year":
			if x.Year != y.Year {
				return x.Year < y.Year
			}
		case "Length":
			if x.Length != y.Length {
				return x.Length < y.Length
			}
		}
	}
	return false
}

func MultiSort(t []*Track, keys []string) sort.Interface { return multiSort{t, keys} }
