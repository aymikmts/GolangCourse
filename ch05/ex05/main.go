// Ex05 は標準入力から得たurlをGETリクエストし、
// そのドキュメント内に含まれる単語と画像の数を返します。
package main

import (
	"GolangCourse/ch05/ex05/countWordsAndImages"
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		url := input.Text()
		words, images, err := countwordsandimages.CountWordsAndImages(url)

		if err != nil {
			fmt.Fprintf(os.Stderr, "ex05: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("[%s]\nwords: %d\timages:%d\n", url, words, images)
	}
}
