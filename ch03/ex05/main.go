// Ex05はフルカラーのマンデルブロ集合を描画します。
// フルカラー化するには、コマンド引数に"-color"を追加します。
// フルカラーマンデルブロ集合の実装は"ch03/mandelbrot/mundelbrotEx05.go"
package main

import (
	//"GolangCourse/ch03/mandelbrot" // GOPATH以下指定
	"flag"
	"io"
	"os"

	"../mandelbrot" // 相対パス指定
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

	mandelbrot.DrawFractal(output)
}
