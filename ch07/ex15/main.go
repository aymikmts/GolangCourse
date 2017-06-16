// Ex15は標準入力から単一の式を読み込み、その式内の変数に対する値をユーザーに問い合わせて、
// それからその結果の環境のもとでその式を評価します。
package main

import (
	"os"
	"bufio"
	"fmt"
	"./eval"
)

func main(){

	var expr string
	fmt.Println("input expr:")
	input := bufio.NewScanner(os.Stdin)
	if input.Scan(){
		fmt.Println(input.Text())
		expr = input.Text()
	}

	var env eval.Env
	fmt.Println("input vals:")
	input := bufio.NewScanner(os.Stdin)
	if input.Scan(){
		fmt.Println(input.Text())
		val := float64()
		Env{"x": input.Text()}
	}


	expr2, err := eval.Parse(expr)
	if err != nil {
		log.Fatal(err)
	}
}