package comma

import (
	"testing"
)

func TestCommaWithBuffer(t *testing.T) {
	var tests = []struct {
		input string
		want  string
	}{
		{"1", "1"},
		{"12", "12"},
		{"123", "123"},
		{"12345", "12,345"},
		{"1234567", "1,234,567"},
		{"123456789", "123,456,789"},
	}

	for _, test := range tests {
		got := CommaWithBuffer(test.input)
		if got != test.want {
			t.Errorf("CommaWithBuffer(%v) = %v", test.input, got)
		}
	}
}

func TestCommaSignedFloat(t *testing.T) {
	var tests = []struct {
		input string
		want  string
	}{
		{"1.0e+8", "100,000,000"},
		{"-1.0e+8", "-100,000,000"},
		{"1.23456789e+8", "123,456,789"},
		{"-1.23456789e+8", "-123,456,789"},
		{"1.2345e+04", "12,345"},
		{"-1.2345e+04", "-12,345"},
		{"1.2345678e+04", "12,345.678"},
		{"-1.2345678e+04", "-12,345.678"},
		{"1.2345e-01", "0.12345"},
		{"-1.2345e-01", "-0.12345"},
	}

	for _, test := range tests {
		got := CommaSignedFloat(test.input)
		if got != test.want {
			t.Errorf("CommaSignedFloat(%v) = %v", test.input, got)
		}
	}
}
