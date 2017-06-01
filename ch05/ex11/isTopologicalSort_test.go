package main

import (
	"strings"
	"testing"
)

var test_noerr1 = map[string][]string{
	"no5": {"no3"},
	"no4": {"no3"},
	"no3": {"no1"},
	"no2": {"no1"},
}

var test_err1 = map[string][]string{
	"no5": {"no3"},
	"no4": {"no3"},
	"no3": {"no1", "no4"},
	"no2": {"no1"},
}

var test_noerr2 = map[string][]string{
	"no6": {"no1", "no3"},
	"no5": {"no3"},
	"no4": {"no3"},
	"no3": {"no1"},
	"no2": {"no1"},
}

var test_err2 = map[string][]string{
	"no6": {"no1", "no3"},
	"no5": {"no3"},
	"no4": {"no3"},
	"no3": {"no1", "no4"},
	"no2": {"no1"},
	"no1": {"no6"},
}

func TestTopoSort(t *testing.T) {
	var tests = []struct {
		source  map[string][]string
		isCycle bool
	}{
		{test_noerr1, false},
		{test_err1, true},
		{test_noerr2, false},
		{test_err2, true},
	}

	for _, test := range tests {
		_, err := topoSort(test.source)
		if test.isCycle {
			if err == nil || !strings.Contains(err.Error(), "cycle") {
				t.Errorf("input slice is cycled. but got other error: %v\n", err)
			}
		} else {
			if err != nil {
				t.Errorf("input slice is not cycled. but got other error: %v\n", err)
			}
		}
	}
}
