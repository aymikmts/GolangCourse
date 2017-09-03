package json

import (
	gojson "encoding/json"
	"fmt"
	"log"
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

func TestMarshalMovie(t *testing.T) {
	in := Movie{
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

	var goJson Movie
	buf, err := gojson.Marshal(in)
	if err != nil {
		log.Fatalln(err)
	}
	if err = gojson.Unmarshal(buf, &goJson); err != nil {
		log.Fatalln(err)
	}

	var myJson Movie
	buf, err = Marshal(in)
	if err != nil {
		log.Fatalln(err)
	}
	if err = gojson.Unmarshal(buf, &myJson); err != nil {
		log.Fatalln(err)
	}

	if !reflect.DeepEqual(myJson, goJson) {
		t.Errorf("my json is not equal to the result of encode/json package.\nin:\n%v\nencode/json package:\n%v",
			myJson, goJson)
	}
}

type TestStruct struct {
	IntVal   int
	Str      string
	Ptr      *int
	StrSlice []string
	StrMap   map[string]string
	BoolVal  bool
	F32Val   float32
	F64Val   float64
}

var val int = 10

func TestMarshal(t *testing.T) {
	tests := []struct {
		s TestStruct
	}{
		{TestStruct{IntVal: 100}},
		{TestStruct{Str: "test"}},
		{TestStruct{Ptr: &val}},
		{TestStruct{StrSlice: []string{"foo", "bar"}}},
		{TestStruct{StrMap: map[string]string{"1": "foo", "2": "bar"}}},
		{TestStruct{BoolVal: true}},
		{TestStruct{F32Val: 0.125}},
		{TestStruct{F64Val: 0.125}},
	}

	for _, test := range tests {

		var want TestStruct
		wantbuf, err := gojson.Marshal(test.s)
		if err = gojson.Unmarshal(wantbuf, &want); err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		var got TestStruct
		gotbuf, err := Marshal(test.s)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		if err = gojson.Unmarshal(gotbuf, &got); err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("\ngot:\n\t%v\n\t%v\nwant:\n\t%v\n\t%v.", got, string(gotbuf), want, string(wantbuf))
		}
	}
}
