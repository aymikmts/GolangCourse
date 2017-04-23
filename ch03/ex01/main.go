// Ex01は3-D面の関数のSVGレンダリングを計算します。
// fの値がNaNだった場合はStderr出力し, そのセルについてはxmlファイル出力をスキップします。
// 実装は"ch03/surface/surface.go"
package main

import (
	"io"
	"os"

	"../surface" // 相対パス指定
	//	"GolangCourse/ch03/surface	// GOPATH以下指定"
)

var output io.Writer = os.Stdout // mainで使用

func main() {
	surface.PrintXML(output)
}
