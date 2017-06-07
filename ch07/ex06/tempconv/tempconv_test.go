package tempconv

import "testing"

func TestCToK(t *testing.T) {
	var tests = []struct {
		input Celsius
		want  Kelvin
	}{
		{Celsius(0), Kelvin(273.15)},
		{Celsius(-273.15), Kelvin(0)},
	}
	for _, test := range tests {
		if got := CToK(test.input); got != test.want {
			t.Errorf(`CToK(%v) = %v`, test.input, got)
		}
	}
}

func TestFToK(t *testing.T) {
	var tests = []struct {
		input Fahrenheit
		want  Kelvin
	}{
		{Fahrenheit(32), Kelvin(273.15)},
		{Fahrenheit(100), Kelvin(310.928)},
	}
	for _, test := range tests {
		if got := FToK(test.input); got != test.want {
			t.Errorf(`FToK(%v) = %v`, test.input, got)
		}
	}
}

func TestKToC(t *testing.T) {
	var tests = []struct {
		input Kelvin
		want  Celsius
	}{
		{Kelvin(273.15), Celsius(0)},
		{Kelvin(0), Celsius(-273.15)},
	}
	for _, test := range tests {
		if got := KToC(test.input); got != test.want {
			t.Errorf(`KToC(%v) = %v`, test.input, got)
		}
	}
}

func TestKToF(t *testing.T) {
	var tests = []struct {
		input Kelvin
		want  Fahrenheit
	}{
		{Kelvin(273.15), Fahrenheit(32)},
		{Kelvin(310.928), Fahrenheit(100)},
	}
	for _, test := range tests {
		if got := KToF(test.input); got != test.want {
			t.Errorf(`KToF(%v) = %v`, test.input, got)
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
		{float64(123.999999), int(3) ,float64(124.000)},
		{float64(999.999999), int(3) ,float64(1000.000)},
	}
	for _, test := range tests {
		if got := round(test.input, test.place); got != test.want {
			t.Errorf(`round(%v) = %v`, test.input, got)
		}
	}
}
