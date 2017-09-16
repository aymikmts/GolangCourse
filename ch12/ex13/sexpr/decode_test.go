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

// ex13 で追加
func TestUnmarshalFieldTag(t *testing.T) {
	type TestStruct struct {
		IntVal   int `sexpr:"int"`
		Str      string
		StrSlice []string `sexpr:"slice"`
	}

	input := []byte("((int 123456789) (Str \"Hello\") (slice (\"test1\" \"test2\")))")
	want := TestStruct{
		IntVal:   123456789,
		Str:      "Hello",
		StrSlice: []string{"test1", "test2"},
	}

	var got TestStruct
	err := Unmarshal(input, &got)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Unmarshal() got \"%v\", want is \"%v\"", got, want)
	}
}

// ex10 で追加
func TestUnmarshal(t *testing.T) {
	type TestStruct struct {
		IntVal    interface{}
		UIntVal   interface{}
		Str       interface{}
		StrSlice  interface{}
		StrArr    interface{}
		StrMap    interface{}
		BoolVal   interface{}
		F32Val    interface{}
		F64Val    interface{}
		Cmp64Val  interface{}
		Cmp128Val interface{}
	}

	input := TestStruct{
		IntVal:    int(10),
		UIntVal:   uint(20),
		Str:       "test",
		StrSlice:  []string{"test1", "test2"},
		StrArr:    [3]string{"test1", "test2", "test3"},
		StrMap:    map[string]string{"a": "A", "b": "B"},
		BoolVal:   true,
		F32Val:    float32(0.125),
		F64Val:    float64(0.125),
		Cmp64Val:  complex64(1.0 + 2.0i),
		Cmp128Val: complex128(3.0 + 4.0i),
	}

	buf, err := Marshal(input)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}
	//fmt.Printf("Marshal() = %s\n", buf)

	var got TestStruct
	if err := Unmarshal(buf, &got); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	//fmt.Printf("Unmarshal() = %v\n", got)

	if !reflect.DeepEqual(input, got) {
		t.Fatalf("input is not equal to got.\nin:\n%v\ngot:\n%v", input, got)
	}
}

// ex09 で追加
func TestToken(t *testing.T) {
	tests := []struct {
		in     string
		tokens []Token
	}{
		{"", []Token{}},
		{"()", []Token{StartList{}, EndList{}}},
		{"(a)", []Token{
			StartList{},
			Symbol{Name: "a"},
			EndList{}},
		},
		{"(\"a\")", []Token{
			StartList{},
			String{Value: "a"},
			EndList{}},
		},
		{"(12345)", []Token{
			StartList{},
			Int{Value: 12345},
			EndList{}},
		},
		{"(a b)", []Token{
			StartList{},
			Symbol{Name: "a"},
			Symbol{Name: "b"},
			EndList{}},
		},
		{"(a \"b\")", []Token{
			StartList{},
			Symbol{Name: "a"},
			String{Value: "b"},
			EndList{}},
		},
		{"(a 12345)", []Token{
			StartList{},
			Symbol{Name: "a"},
			Int{Value: 12345},
			EndList{}},
		},
		{"((a 1) (b \"test\"))", []Token{
			StartList{},
			StartList{}, Symbol{Name: "a"}, Int{Value: 1}, EndList{},
			StartList{}, Symbol{Name: "b"}, String{Value: "test"}, EndList{},
			EndList{}},
		},
	}

	for _, test := range tests {
		d := NewDecoder(bytes.NewReader([]byte(test.in)))
		for _, want := range test.tokens {
			got := d.Token()
			if !reflect.DeepEqual(got, want) {
				t.Errorf("got is %#v, want is %#v", got, want)
			}
		}
	}
}

// ex08 で追加
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
