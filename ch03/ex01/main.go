// Ex01は3-D面の関数のSVGレンダリングを計算します。
// fの値がNaNだった場合はStderr出力し, そのセルについてはxmlファイル出力をスキップします。
// 実装は"ch03/surface/surface.go"
package main

import (
	"GolangCourse/ch03/surface"
	"io"
	"os"
)

var output io.Writer = os.Stdout // mainで使用

func main() {
	surface.PrintXML(output)
}
