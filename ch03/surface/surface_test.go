package surface

import (
	"bytes"
	"testing"
)

func TestPrintXML(t *testing.T) {
	out := new(bytes.Buffer)
	PrintXML(out)

	// バッファから文字列"NaN"を探索し、ヒットしたら失敗。
	sep := []byte("NaN")
	buf := out.Bytes()
	index := bytes.Index(buf, sep)
	if index != -1 {
		t.Errorf(`"NaN" has hit. index=%v`, index)
	}
}

func TestCorner(t *testing.T) {
	var tests = []struct {
		i  int
		j  int
		ok bool
	}{
		{0, 0, true},
		{50, 50, false},
	}
	for _, test := range tests {
		if _, _, got := corner(test.i, test.j); got != test.ok {
			t.Errorf(`corner(%v, %v) = %v`, test.i, test.j, got)
		}
	}
}

func TestCalcColor(t *testing.T) {
	var tests = []struct {
		input float64
		want  string
	}{
		{0, "#0000ff"},
		{127, "#7f0080"},
		{128, "#80007f"},
		{255, "#ff0000"},
	}
	for _, test := range tests {
		zMin = 0
		zMax = 255
		if got := calcColor(test.input); got != test.want {
			t.Errorf(`calcColor(%v) = %v`, test.input, got)
		}
	}
}
