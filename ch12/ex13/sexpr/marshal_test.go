package sexpr

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

// ex13 で追加
func TestMarshalFieldTag(t *testing.T) {
	type TestStruct struct {
		IntVal   int `sexpr:"int"`
		Str      string
		StrSlice []string `sexpr:"slice"`
	}
	input := TestStruct{
		IntVal:   123456789,
		Str:      "Hello",
		StrSlice: []string{"test1", "test2"},
	}
	want := "((int 123456789) (Str \"Hello\") (slice (\"test1\" \"test2\")))"

	got, err := Marshal(input)
	if err != nil {
		t.Fatal(err)
	}
	if string(got) != want {
		t.Errorf("Marshal() got \"%s\", want is \"%s\"", got, want)
	}
}

func TestMarshalIgnoreZeroValue(t *testing.T) {
	type TestStruct struct {
		IntVal    int
		Str       string
		Ptr       *int
		StrSlice  []string
		StrMap    map[string]string
		BoolVal   bool
		F32Val    float32
		F64Val    float64
		Cmp64Val  complex64
		Cmp128Val complex128
		Interface interface{}
	}

	tests := []struct {
		in   TestStruct
		want []byte
	}{
		{TestStruct{}, []byte("")},
		{TestStruct{IntVal: 0}, []byte("")},
		{TestStruct{Str: ""}, []byte("")},
		{TestStruct{Ptr: nil, Interface: nil}, []byte("")},
		{TestStruct{StrSlice: []string{}}, []byte("")},
		{TestStruct{StrMap: map[string]string{}}, []byte("")},
		{TestStruct{BoolVal: false}, []byte("")},
		{TestStruct{F32Val: 0, F64Val: 0}, []byte("")},
		{TestStruct{Cmp64Val: 0, Cmp128Val: 0}, []byte("")},

		{TestStruct{IntVal: 1, Str: "test"}, []byte("((IntVal 1) (Str \"test\"))")},
		{TestStruct{StrSlice: []string{"a", "b"}}, []byte("((StrSlice (\"a\" \"b\")))")},
	}

	for _, test := range tests {
		got, err := Marshal(test.in)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}

		if !bytes.Equal(got, test.want) {
			t.Errorf("Marshal(%v)\ngot:\n\"%s\"\nwant:\n\"%s\"",
				test.in, string(got), string(test.want))
		}
	}
}
