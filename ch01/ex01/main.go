// mainは、起動したコマンド名と引数に与えられた値を出力します。
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
	s, sep := "", ""
	for _, arg := range args[:] {
		s += sep + arg
		sep = " "
	}
	fmt.Fprintln(out, s)
	return nil
}
