// Ex08はcomplex64, complex128, big.Float, big.Ratの4つの異なる数値の表現を使って同じフラクタルを実装します。
// Stderrに各条件で画像出力する前後のメモリ量を出力します。
// 各条件は、コマンド引数に"-format [format type]"を追加します。
// 実装は"ch03/mandelbrot/mundelbrotEx08.go"
package main

import (
	//"GolangCourse/ch03/mandelbrot" // GOPATH以下指定
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"runtime"

	"../mandelbrot" // 相対パス指定
)

var output io.Writer = &bytes.Buffer{}
var mem runtime.MemStats

// 精度をコマンド引数で指定
var formatFlag = flag.String("format", "default", "format type: default/complex64/complex128/big.Float/big.Rat")

// メモリ量をstderrに表示する
func printMemory() {
	runtime.ReadMemStats(&mem)
	fmt.Fprintf(os.Stderr, "total memory: %v\n", mem.TotalAlloc)
}

func main() {
	// メモリ量を表示する
	printMemory()

	// 計算時間がかかるため、iteration値を変更する
	mandelbrot.IterationsVal = 1

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

	// メモリ量を表示する
	printMemory()
}
