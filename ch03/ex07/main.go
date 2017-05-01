// Ex07はニュートン法を用いたフラクタル画像を描画します。
// ニュートン法を用いたフラクタル画像を描画するには、
// コマンド引数に"-fractal newton"を追加します。
// 実装は"ch03/mandelbrot/mundelbrotEx07.go"
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

// 描画するフラクタルをコマンド引数で指定
var fractalFlag = flag.String("fractal", "mandelbrot", "FractalType: mandelbrot/newton")

func main() {
	// flagでモノクロ・カラーを切り替える
	flag.Parse()
	if *colorFlag {
		mandelbrot.IsColoring = true
	}

	switch *fractalFlag {
	case "mandelbrot":
		mandelbrot.Fractal = mandelbrot.MANDELBROT
	case "newton":
		mandelbrot.Fractal = mandelbrot.NEWTON
	default:
		mandelbrot.Fractal = mandelbrot.MANDELBROT

	}
	mandelbrot.DrawFractal(output)
}
