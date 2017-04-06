// echo2は、引数に与えられた値をfor文により代入を繰り返して整形し、表示します。
package echo2

import "fmt"

func Echo(args []string) {
	s, sep := "", ""
	for _, arg := range args[:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
