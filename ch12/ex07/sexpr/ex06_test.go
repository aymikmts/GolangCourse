package sexpr

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

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

func TestMarshalIgnoreZeroValue(t *testing.T) {
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
