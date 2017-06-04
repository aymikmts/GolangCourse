// Ex15は、可変個引数および少なくとも1つの引数が必要なmax, minを実装します。
package main

import (
	"fmt"
)

func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

// maxWithErrは、可変個引数対応のmax関数です。
// 引数なしはエラーとなります。
func maxWithErr(vals ...int) (int, error) {
	if len(vals) <= 0 {
		return 0, fmt.Errorf("maxWithErr() needs more than one param.")
	}

	max := vals[0]
	for _, val := range vals[1:] {
		if val > max {
			max = val
		}
	}
	return max, nil
}

// maxは、少なくとも1つの引数が必要なmax関数です。
func max(first int, vals ...int) int {
	max := first
	for _, val := range vals {
		if val > max {
			max = val
		}
	}
	return max
}

// minWithErrは、可変個引数対応のmin関数です。
// 引数なしはエラーとなります。
func minWithErr(vals ...int) (int, error) {
	if len(vals) <= 0 {
		return 0, fmt.Errorf("maxWithErr() needs more than one param.")
	}

	min := vals[0]
	for _, val := range vals[1:] {
		if val < min {
			min = val
		}
	}
	return min, nil
}

// minは、少なくとも1つの引数が必要なmin関数です。
func min(first int, vals ...int) int {
	min := first
	for _, val := range vals {
		if val < min {
			min = val
		}
	}
	return min
}

func main() {
	fmt.Printf("-- MAX --\n")
	fmt.Println(max(0))
	fmt.Println(maxWithErr())
	fmt.Println(maxWithErr(1, 2, 3, 4))

	fmt.Printf("-- MIN --\n")
	fmt.Println(min(0))
	fmt.Println(minWithErr())
	fmt.Println(minWithErr(1, 2, 3, 4))
}
