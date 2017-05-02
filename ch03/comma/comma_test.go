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
