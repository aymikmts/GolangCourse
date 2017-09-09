package sexpr

import (
	"bytes"
	"fmt"
	"os"
	"reflect"
	"testing"
)

type Movie struct {
	Title, Subtitle string
	Year            int
	Color           bool
	Actor           map[string]string
	Oscars          []string
	Sequel          *string
}

func TestNewDecoder(t *testing.T) {
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

	buf, err := Marshal(strangelove)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	var got, want Movie
	if err = Unmarshal(buf, &got); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	if err = NewDecoder(bytes.NewReader(buf)).Decode(&want); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got is not equal to want.\ngot:\n%v\nwant:\n%v",
			got, want)
	}

}
