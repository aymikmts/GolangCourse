// Ex03は3-D面の関数のSVGレンダリングを計算します。
// コマンド引数フラグによって、描画するモデルに色をつけられます。
// 着色の色設定の実装については"ch03/surface/surfaceEx03.go"
package main

import (
	"GolangCourse/ch03/surface"
	"flag"
	"io"
	"os"
)

var output io.Writer = os.Stdout

// 描画モデルをコマンド引数で指定
var modelFlag = flag.String("model", "Default",
	"Switch drawing model.\n\tmodel are:\"Default\", \"EggCase\", \"Moguls\"")

// 描画モデルに色を付けるかどうかを指定
var modelColor = flag.Bool("gradient", false, "Gradient color ON/OFF")

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

	// flagで着色有無を指定
	if *modelColor {
		surface.IsGradientColor = true
	}

	surface.PrintXML(output)
}
