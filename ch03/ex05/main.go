// Ex05はフルカラーのマンデルブロ集合を描画します。
package main

import (
	"GolangCourse/ch03/mandelbrot" // GOPATH以下指定
	"io"
	"os"
	//"../mandelbrot"	// 相対パス指定
	"flag"
)

var output io.Writer = os.Stdout

// モノクロ・カラーをコマンド引数で指定
var colorFlag = flag.Bool("color", false, "Full color ON/OFF")

func main() {
	// flagでモノクロ・カラーを切り替える
	flag.Parse()
	if *colorFlag {
		mandelbrot.IsColoring = true
	}

	mandelbrot.DrawMandelbrot(output)
}
