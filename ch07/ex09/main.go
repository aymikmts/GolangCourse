package main

import (
	"GolangCourse/ch07/sorting"
	"html/template"
	"log"
	"os"
)

const tmpl = `
<h1>Track List</h1>
<table>
<tr style='text-align: left'>
	<th>Title</th>
	<th>Artist</th>
	<th>Album</th>
	<th>Year</th>
	<th>Length</th>
</tr>

{{range .Tracks}}
<tr>
	<td>{{.Title}}</td>
	<td>{{.Artist}}</td>
	<td>{{.Album}}</td>
	<td>{{.Year}}</td>
	<td>{{.Length}}</td>
</tr>
{{end}}
</table>
`

type TrackListResult struct {
	Tracks []*sorting.Track
}

var tracks = []*sorting.Track{
	{"Go", "Delilah", "From the Roots Up", 2012, sorting.Length("3m38s")},
	{"Go", "Moby", "Moby", 1992, sorting.Length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, sorting.Length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, sorting.Length("4m24s")},
}

func title(keys []string) []string {
	keys = append(keys, "Tracks")
	sorting.MultiSort(tracks, keys)
	return keys
}

func main() {
	var trackList = template.Must(template.New("trackList").Parse(tmpl))
	if err := trackList.Execute(os.Stdout, TrackListResult{tracks}); err != nil {
		log.Fatal(err)
	}

}
