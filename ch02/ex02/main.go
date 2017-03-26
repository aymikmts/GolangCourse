// cfは、その数値引数を摂氏と華氏へ変換します。
package main

import (
	"fmt"
	"os"
	"strconv"

	"../ex01/tempconv"
	"./lengthconv"
	"./weightconv"
)

func main() {
	if len(os.Args) < 2 {
		// コマンドライン引数が指定されなかった場合には標準入力から数値を読み込み、各単位変換結果を表示する
		var arg string
		fmt.Println("Input value:")
		fmt.Scanf("%s", &arg)
		printConvResults(arg)
	} else {
		// コマンドライン引数で指定された数値を読み込み、各単位変換結果を表示する
		for _, arg := range os.Args[1:] {
			printConvResults(arg)
		}
	}
}

// printConvResultsは引数で指定された数値を各単位変換し、結果を表示する。
func printConvResults(arg string) {
	fmt.Printf("--------------\narg = %s\n", arg)
	t, err := strconv.ParseFloat(arg, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cf: %v\n", err)
		os.Exit(1)
	}
	printTemp(t)
	printLength(t)
	printWeight(t)
}

// printTempは引数で指定された数値を温度単位(℃, ℉)に変換し、表示する。
func printTemp(t float64) {
	f := tempconv.Fahrenheit(t)
	c := tempconv.Celsius(t)
	fmt.Printf("Temp\t%s = %s, %s = %s\n",
		f, tempconv.FToC(f), c, tempconv.CToF(c))
}

// printLengthは引数で指定された数値を長さ単位(ポンド、メートル)に変換し、表示する。
func printLength(t float64) {
	f := lengthconv.Feet(t)
	m := lengthconv.Meter(t)
	fmt.Printf("Length\t%s = %s, %s = %s\n",
		f, lengthconv.FToM(f), m, lengthconv.MToF(m))
}

// printWeightは引数で指定された数値を重さ単位(ポンド、キログラム)に変換し、表示する。
func printWeight(t float64) {
	p := weightconv.Pound(t)
	kg := weightconv.KiloGram(t)
	fmt.Printf("Weight\t%s = %s, %s = %s\n",
		p, weightconv.PToKg(p), kg, weightconv.KgToP(kg))
}
