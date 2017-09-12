package params

import "testing"

func TestPack(t *testing.T) {
	type Data struct {
		Labels     []string `http:"l"`
		MaxResults int      `http:"max"`
		Exact      bool     `http:"x"`
	}

	tests := []struct {
		input Data
		want  string
	}{
		{Data{Labels: []string{}, MaxResults: 10, Exact: false}, ""},
		{Data{
			Labels:     []string{"golang", "programming"},
			MaxResults: 10,
			Exact:      false},
			"l=golang&l=programming",
		},
		{Data{
			Labels:     []string{"golang", "programming"},
			MaxResults: 100,
			Exact:      false},
			"l=golang&l=programming&max=100",
		},
		// {Data{
		// 	Labels:     []string{"golang", "programming"},
		// 	MaxResults: 10,
		// 	Exact:      true},
		// 	"l=golang&l=programming&x=true",
		// },
	}

	for _, test := range tests {
		got, err := Pack(test.input)
		if err != nil {
			t.Errorf("Pack(%v) failed: %v", test.input, err)
			continue
		}

		if got != test.want {
			t.Errorf("Pack(%v) got \"%s\", but want is \"%s\".", test.input, got, test.want)
		}
	}
}
