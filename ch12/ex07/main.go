package main

import (
	"GolangCourse/ch12/ex07/sexpr"
	"bytes"
	"fmt"
	"log"
)

type Movie struct {
	Title, Subtitle string
	Year            int
	Color           bool
	Actor           map[string]string
	Oscars          []string
	Sequel          *string
}

func main() {
	strangelove := Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
		Year:     1964,
		Color:    false,
		Actor: map[string]string{
			"Dr. Strangelove":            "Peter Sellers",
			"Grp. Capt. Lionel Mandrake": "Peter Sellers",
			"Pres. Merkin Muffley":       "Peter Sellers",
			"Gen. Buck Turgidson":        "George C. Scott",
			"Brig. Gen. Jack D. Ripper":  "Sterling Hayden",
			`Maj. T.J. "King" Kong`:      "Slim Pickens",
		},
		Oscars: []string{
			"Best Actor (Nomin.)",
			"Best Adapted Screenplay (Nomin.)",
			"Best Director (Nomin.)",
			"Best Picture (Nomin.)",
		},
	}

	buf, err := sexpr.Marshal(strangelove)
	if err != nil {
		log.Fatalln(err)
	}

	r := bytes.NewReader(buf)
	var result Movie
	if err = sexpr.NewDecoder(r).Decode(&result); err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("\n--- RESULT ---\n")
	fmt.Printf("title: %s\nsubtitle: %s\n", result.Title, result.Subtitle)
	fmt.Printf("year: %d\ncolor: %v\n", result.Year, result.Color)
	fmt.Printf("actor:\n")
	for key, val := range result.Actor {
		fmt.Printf("\t%s: %s\n", key, val)
	}
	fmt.Printf("oscars:\n")
	for _, s := range result.Oscars {
		fmt.Printf("\t%s\n", s)
	}
	fmt.Printf("sequel: %v\n", result.Sequel)
}
