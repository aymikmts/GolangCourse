// mainは入力に2回以上現れた行の数とテキスト、その行が含まれていた全てのファイル名を表示します。
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	fnames := make(map[string]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()

			for line, n := range counts {
				if n > 0 {
					fnames[line] += arg + ","
				}
			}
		}
	}
	for line, n := range counts {
		if n > 0 {
			fmt.Printf("%d\t%s\t%s\n", n, line, fnames[line])
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
