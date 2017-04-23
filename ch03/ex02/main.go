// Ex02は3-D面の関数のSVGレンダリングを計算します。
// コマンド引数フラグによって、描画するモデルに色をつけられます。
// 各モデルの実装は"ch03/surface/surfaceEx02.go"
package main

import (
	"flag"
	"io"
	"os"

	"../surface" // 相対パス指定
	//"GolangCourse/ch03/surface"	// GOPATH以下指定
)

var output io.Writer = os.Stdout

// 描画モデルをコマンド引数で指定
var modelFlag = flag.String("model", "Default",
	"Switch drawing model.\n\tmodel are:\"Default\", \"EggCase\", \"Moguls\"")

func main() {
	// flagで描画モデルを切り替える。
	flag.Parse()
	switch *modelFlag {
	case "Default":
		surface.Model = surface.DEFAULT
	case "EggCase":
		surface.Model = surface.EGGCASE
	case "Moguls":
		surface.Model = surface.MOGULS
	case "Saddle":
		surface.Model = surface.SADDLE
	default:
		surface.Model = surface.DEFAULT
	}
	surface.PrintXML(output)
}
