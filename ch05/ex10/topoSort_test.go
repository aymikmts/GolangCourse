package main

import (
	"testing"
)

var test1 = map[string]map[string]bool{
	"no5": {"no3": true},
	"no4": {"no3": true},
	"no3": {"no1": true},
	"no2": {"no1": true},
}

func TestIsToporicalSort(t *testing.T) {
	var tests = []struct {
		input map[string]map[string]bool
		ts    []string
		want  bool
	}{
		{test1, []string{"no1", "no2", "no3", "no4", "no5"}, true},
		{test1, []string{"no1", "no2", "no3", "no5", "no4"}, true},
		{test1, []string{"no1", "no3", "no2", "no4", "no5"}, true},
		{test1, []string{"no1", "no3", "no4", "no5", "no2"}, true},
		{test1, []string{"no1", "no3", "no5", "no4", "no2"}, true},

		{test1, []string{"no1", "no2", "no4", "no3", "no5"}, false},
		{test1, []string{"no1", "no2", "no4", "no5", "no5"}, false},
		{test1, []string{"no1", "no2", "no5", "no3", "no4"}, false},
		{test1, []string{"no1", "no5", "no4", "no3", "no2"}, false},
		{test1, []string{"no5", "no4", "no3", "no2", "no1"}, false},
	}

	for _, test := range tests {
		prereqs = test.input
		got := isTopologicalSorted(test.ts)
		if got != test.want {
			t.Errorf("want is %v, but got is %v. topo sort: %v\n", test.want, got, test.ts)
		}
	}
}
