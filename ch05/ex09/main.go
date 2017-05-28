// Ex09は、文字列s内の部分文字列$fooをf("foo")によって置換します。
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		input := scanner.Text()
		fmt.Printf("----------------------\ninput: %s\n", input)
		out := expand(input, toUpper)
		fmt.Printf("outpt: %s\n", out)
	}
}

func expand(s string, f func(string) string) string {
	if f == nil {
		return s
	}

	slice := strings.Split(s, " ")
	for _, word := range slice {
		if strings.HasPrefix(word, "$") {
			dst := f(strings.TrimPrefix(word, "$"))
			s = strings.Replace(s, word, dst, 1)
		}
	}

	return s
}

func toUpper(s string) string {
	return strings.ToUpper(s)
}
