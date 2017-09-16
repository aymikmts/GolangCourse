package main

import (
	"GolangCourse/ch12/ex13/sexpr"
	"fmt"
	"log"
)

type Movie struct {
	Title    string            `sexpr:"t"`
	Subtitle string            `sexpr:"sub"`
	Year     int               `sexpr:"y"`
	Color    bool              `sexpr:"c"`
	Actor    map[string]string `sexpr:"a"`
	Oscars   []string          `sexpr:"o"`
	Sequel   *string           `sexpr:"s"`
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

	fmt.Println("--- Show S-expr Marshal ---")
	buf, err := sexpr.Marshal(strangelove)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(buf))

	/////////////////////////////
	fmt.Println("--- Unmarshal ---")
	var got Movie
	err = sexpr.Unmarshal(buf, &got)
	fmt.Println(got)
}
