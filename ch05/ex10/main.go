// Ex10は、スライスの代わりにマップを使ったtopoSortです。
package main

import (
	"fmt"
)

var prereqs = map[string]map[string]bool{
	"algorithms": {"data structures": true},
	"calculus":   {"liner algebra": true},
	"compilers": {
		"data structures":       true,
		"formal languages":      true,
		"computer organization": true,
	},

	"data structures":       {"discrete math": true},
	"databases":             {"data structures": true},
	"discrete math":         {"intro to programming": true},
	"formal languages":      {"discrete math": true},
	"networks":              {"operating systems": true},
	"operating systems":     {"data structures": true, "computer organization": true},
	"programming languages": {"data structures": true, "computer organization": true},
}

func main() {
	courses := topoSort(prereqs)

	for i, course := range courses {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}

	fmt.Printf("isToporicalSorted: %v\n", isTopologicalSorted(courses))
}

// topoSortは、深さ優先探索して正当な順序を計算するトポロジカルソートです。
func topoSort(m map[string]map[string]bool) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items map[string]bool)

	visitAll = func(items map[string]bool) {
		for item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}

	keys := make(map[string]bool)
	for key := range m {
		keys[key] = true
	}
	fmt.Println(keys)
	visitAll(keys)
	return order
}

// isTopologicalSortedは、スライスがトポロジカル順序になっているかどうかを返します。
func isTopologicalSorted(input []string) bool {
	node := make(map[string]bool)

	for i := 0; i < len(input); i++ {
		node[input[i]] = true
		if i == 0 {
			continue
		}

		for key, _ := range prereqs[input[i]] {
			if node[key] == false {
				return false
			}
		}
	}

	return true
}
