// mainは、個々の引数のインデックスと値の組を1行ごとに出力します。
package main

import (
	"fmt"
	"io"
	"os"
)

var out io.Writer = os.Stdout

func main() {
	if err := echo(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "echo: %v\n", err)
		os.Exit(1)
	}
}

func echo(args []string) error {
	for i := 1; i < len(args); i++ {
		fmt.Fprintln(out, i, args[i])
	}
	return nil
}
