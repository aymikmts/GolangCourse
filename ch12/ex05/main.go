package main

import (
	ex05json "GolangCourse/ch12/ex05/json"
	json "encoding/json"
	"fmt"
	"log"
	"reflect"
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

	fmt.Println("-- encode/json --")
	buf, err := json.Marshal(strangelove)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(buf))
	var goJson Movie
	if err = json.Unmarshal(buf, &goJson); err != nil {
		log.Fatalln(err)
	}

	fmt.Println("\n-- my json package --")
	buf, err = ex05json.Marshal(strangelove)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(buf))
	var ex05Json Movie
	if err = json.Unmarshal(buf, &ex05Json); err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("\nreflect.DeepEqual(goJson, ex05Json): %v\n", reflect.DeepEqual(goJson, ex05Json))

}
