package fetchall

import "testing"

func TestAddPrefix(t *testing.T) {
	var tests = []struct {
		input  string
		output string
	}{
		{"test", "http://test"},
	}

	for _, test := range tests {
		if got := addprefix(test.input); got != test.output {
			t.Errorf("addprefix(%v)= %v", test.input, got)
		}
	}
}
