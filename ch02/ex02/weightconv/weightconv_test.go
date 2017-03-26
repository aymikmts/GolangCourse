package weightconv

import "testing"

func TestPToKg(t *testing.T) {
	var tests = []struct {
		input Pound
		want  KiloGram
	}{
		{Pound(0), KiloGram(0)},
		{Pound(1), KiloGram(0.454)},
		{Pound(100), KiloGram(45.359)},
	}
	for _, test := range tests {
		if got := PToKg(test.input); got != test.want {
			t.Errorf(`PToKg(%v) = %v`, test.input, got)
		}
	}
}

func TestKgToP(t *testing.T) {
	var tests = []struct {
		input KiloGram
		want  Pound
	}{
		{KiloGram(0), Pound(0)},
		{KiloGram(1), Pound(2.205)},
		{KiloGram(100), Pound(220.462)},
	}
	for _, test := range tests {
		if got := KgToP(test.input); got != test.want {
			t.Errorf(`KgToP(%v) = %v`, test.input, got)
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
