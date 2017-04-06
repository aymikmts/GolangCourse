// echo3は、引数に与えられた値をstringパッケージを使って整形し、表示します。
package echo3

import (
	"fmt"
	"strings"
)

func Echo(args []string) {
	s := strings.Join(args[:], " ")
	fmt.Println(s)
}
