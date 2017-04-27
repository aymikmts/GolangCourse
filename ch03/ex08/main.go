// Ex08は
// 実装は"ch03/mandelbrot/mundelbrotEx08.go"
package main

import (
	//"GolangCourse/ch03/mandelbrot" // GOPATH以下指定
	"flag"
	"io"
	"os"

	"../mandelbrot" // 相対パス指定
)

var output io.Writer = os.Stdout

// 精度をコマンド引数で指定
var formatFlag = flag.String("format", "default", "format type: default/complex64/complex128/big.Float/big.Rat")

func main() {
	// 計算時間がかかるため、iteration値を変更する
	mandelbrot.IterationsVal = 200

	// flagでフラクタル精度を切り替える
	flag.Parse()
	switch *formatFlag {
	case "complex64":
		mandelbrot.FormatCompTest = mandelbrot.CMPLX64
	case "complex128":
		mandelbrot.FormatCompTest = mandelbrot.CMPLX128
	case "big.Float":
		mandelbrot.FormatCompTest = mandelbrot.BIGFLOAT
	case "big.Rat":
		mandelbrot.FormatCompTest = mandelbrot.BIGRAT
	default:
		mandelbrot.FormatCompTest = mandelbrot.DEFAULT
	}
	mandelbrot.DrawFractal(output)
}
