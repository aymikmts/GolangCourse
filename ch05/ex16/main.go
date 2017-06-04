// Ex16は、strings.Joinの可変長引数バージョンです。
package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// joinByVariableは、strings.Joinの可変長引数バージョンです。
// 引数が2つ以下であればエラーとなります。
func joinByVariable(in ...string) (string, error) {
	if len(in) < 2 {
		return "", fmt.Errorf("input num is %d, but want is more than 2.", len(in))
	}
	s := strings.Join(in[:len(in)-1], in[len(in)-1])
	return s, nil
}

func main() {
	s, err := joinByVariable(os.Args[1:]...)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(s)
}
