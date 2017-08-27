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

func ExampleMarshal() {
	val := 987654321
	testStruct := TestStruct{
		IntVal: 123456789,
		Str:    "Hello, World",
		Ptr:    &val,
		StrSlice: []string{
			"str1", "str2", "str3",
		},
		StrMap: map[string]string{
			"key1": "val1",
			"key2": "val2",
			"key3": "val3",
		},
		BoolVal:   true,
		F32Val:    0.1,
		F64Val:    0.1,
		Cmp64Val:  complex(float32(1.0), float32(2.0)),
		Cmp128Val: complex(float64(1.0), float64(2.0)),
		Interface: []int{10, 20, 30},
	}

	buf, err := Marshal(testStruct)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	fmt.Println(string(buf))

	// Output:
	// ((IntVal 123456789) (Str "Hello, World") (Ptr 987654321) (StrSlice ("str1" "str2" "str3")) (StrMap (("key1" "val1") ("key2" "val2") ("key3" "val3"))) (BoolVal t) (F32Val 0.100000) (F64Val 0.100000) (Cmp64Val #C(1.000000 2.000000)) (Cmp128Val #C(1.000000 2.000000)) (Interface ("[]int" (10 20 30))))
	//
}

func TestMarshal(t *testing.T) {
	tests := []struct {
		s    TestStruct
		want []byte
	}{
		{TestStruct{BoolVal: true}, []byte("(BoolVal t)")},
		{TestStruct{BoolVal: false}, []byte("(BoolVal nil)")},
		{TestStruct{Interface: nil}, []byte("(Interface nil)")},
		{TestStruct{Interface: int(123456789)}, []byte("(Interface (\"int\" 123456789))")},
	}

	for _, test := range tests {
		got, err := Marshal(test.s)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		if !bytes.Contains(got, test.want) {
			t.Errorf("got \"%s\" is not contain \"%s\"", string(got), string(test.want))
		}
	}
}
