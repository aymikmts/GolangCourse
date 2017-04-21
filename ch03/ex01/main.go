// Ex01は3-D面の関数のSVGレンダリングを計算します。
// fの値がNaNだった場合はStderr出力し、xmlファイル出力をスキップします。
package main

import (
	"GolangCourse/ch03/surface"
	"io"
	"os"
)

var output io.Writer = os.Stdout // mainで使用

func main() {
	// 実装は"ch03/surface/surface.go"
	surface.PrintXML(output)
}
