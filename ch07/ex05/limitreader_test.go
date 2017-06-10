package main

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func TestLimitReader(t *testing.T) {
	var tests = []struct {
		input []byte
		n     int64
		want  []byte
	}{
		{[]byte(""), 0, []byte("")},
		{[]byte("hello, world!"), 0, []byte("")},
		{[]byte("hello, world!"), 6, []byte("hello,")},
		{[]byte("hello, world!"), int64(len("hello, world!")), []byte("hello, world!")},
	}
	for _, test := range tests {
		r := bytes.NewReader(test.input)
		lmt := LimitReader(r, test.n)
		buf, err := ioutil.ReadAll(lmt)
		if err != nil {
			t.Errorf("%v", err)
		}
		if bytes.Compare(buf, test.want) != 0 {
			t.Errorf("buf: %v\n", string(buf))
		}
	}
}
