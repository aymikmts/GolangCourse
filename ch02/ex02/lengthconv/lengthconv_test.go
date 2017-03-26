package lengthconv

import "testing"

func TestFToM(t *testing.T) {
	var tests = []struct {
		input Feet
		want  Meter
	}{
		{Feet(0), Meter(0)},
		{Feet(1), Meter(0.305)},
		{Feet(100), Meter(30.48)},
	}
	for _, test := range tests {
		if got := FToM(test.input); got != test.want {
			t.Errorf(`FToM(%v) = %v`, test.input, got)
		}
	}
}

func TestMToF(t *testing.T) {
	var tests = []struct {
		input Meter
		want  Feet
	}{
		{Meter(0), Feet(0)},
		{Meter(1), Feet(3.281)},
		{Meter(100), Feet(328.084)},
	}
	for _, test := range tests {
		if got := MToF(test.input); got != test.want {
			t.Errorf(`MToF(%v) = %v`, test.input, got)
		}
	}
}

func TestRound(t *testing.T) {
	var tests = []struct {
		input float64
		place int
		want  float64
	}{
		{float64(123.555555), int(3), float64(123.556)},
		{float64(123.444444), int(3), float64(123.444)},
		{float64(123.999999), int(3), float64(124.000)},
		{float64(999.999999), int(3), float64(1000.000)},
	}
	for _, test := range tests {
		if got := round(test.input, test.place); got != test.want {
			t.Errorf(`round(%v) = %v`, test.input, got)
		}
	}
}
