// Ex11は、循環を報告する機能を付けたtopoSortです。
package main

import (
	"fmt"
	"log"
	"sort"
)

var prereqs = map[string][]string{
	"algorithms":    {"data structures"},
	"calculus":      {"liner algebra"},
	"liner algebra": {"calculus"}, // 循環
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	courses, err := topoSort(prereqs)
	if err != nil {
		log.Fatalln(err)
	}

	for i, course := range courses {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) ([]string, error) {
	var order []string
	seen := make(map[string]bool)

	var visitAll func(items []string)

	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	visitAll(keys)

	err := isTopologicalSorted(order, m)
	if err != nil {
		return nil, err
	}
	return order, nil
}

// isTopologicalSortedは、スライスがトポロジカル順序になっているかどうかを返します。
func isTopologicalSorted(ts []string, source map[string][]string) error {
	node := make(map[string]bool)

	for i := 0; i < len(ts); i++ {
		node[ts[i]] = true
		if i == 0 {
			continue
		}

		for _, key := range source[ts[i]] {
			if !node[key] {
				for _, key2 := range source[key] {
					if node[key2] {
						return fmt.Errorf("\"%v\" and \"%v\" is cycled.\n", key, key2)

					}
				}
				return fmt.Errorf("is not topological sorted.\n")
			}
		}
	}

	return nil
}
