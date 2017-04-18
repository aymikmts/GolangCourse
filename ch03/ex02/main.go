// Ex02は3-D面の関数のSVGレンダリングを計算します。
// コマンド引数フラグによって、描画するモデルを切り替えます。
package main

import (
	"flag"
	"fmt"
	"io"
	"math"
	"os"
)

// 描画するものの種類を定義
type ModelType int

const (
	Default ModelType = iota
	EggCase
	Moguls
	Saddle
)

var modelType ModelType = Default // デフォルトモデル
var output io.Writer = os.Stdout  // mainで使用
var modelFlag = flag.String("model", "Default",
	"Switch drawing model.\n\tmodel are:\"Default\", \"EggCase\", \"Moguls\"")

const (
	width, height = 600, 320            // キャンバスの大きさ(画素数)
	cells         = 100                 // 格子のます目の数
	xyrange       = 30.0                // 軸の範囲 (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // x単位 および y単位あたりの画素数
	zscale        = height * 0.4        // z単位あたりの画素数
	angle         = math.Pi / 6         // x, y軸の角度 (=30度)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30度), cos(30度)

func main() {
	// flagで描画モデルを切り替える。
	flag.Parse()
	switch *modelFlag {
	case "Default":
		modelType = Default
	case "EggCase":
		modelType = EggCase
	case "Moguls":
		modelType = Moguls
	case "Saddle":
		modelType = Saddle
	default:
		modelType = Default
	}
	printXML(output)

}

func printXML(out io.Writer) {
	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			// 戻り値okがfalseだったときは、スキップする。
			ax, ay, ok := corner(i+1, j)
			if !ok {
				continue
			}
			bx, by, ok := corner(i, j)
			if !ok {
				continue
			}
			cx, cy, ok := corner(i, j+1)
			if !ok {
				continue
			}
			dx, dy, ok := corner(i+1, j+1)
			if !ok {
				continue
			}
			fmt.Fprintf(out, "<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintln(out, "</svg>")
}

func corner(i, j int) (sx float64, sy float64, ok bool) {
	// ます目(i,j)のかどの点(x, y)を見つける。
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// 面の高さzを計算する。
	z := f(x, y)
	if math.IsNaN(z) {
		// 値がNaNだった場合はStderr出力し、falseを返す。
		fmt.Fprintf(os.Stderr, "[%d, %d]z := f(%.5f, %.5f) : Value is NaN!\n", i, j, x, y)
		ok = false
		return
	}

	// (x, y, z)を2-D SVGキャンバス(sx, sy)へ等角的に投影。
	sx = width/2 + (x-y)*cos30*xyscale
	sy = height/2 + (x+y)*sin30*xyscale - z*zscale
	ok = true
	return
}

// modelTypeによって描画するモデルをswitch
func f(x, y float64) (r float64) {
	switch modelType {
	case Default:
		r = fDefault(x, y)

	case EggCase:
		r = fEggCase(x, y)

	case Moguls:
		r = fMoguls(x, y)

	case Saddle:
		r = fSaddle(x, y)

	default:
		r = fDefault(x, y)
	}
	return
}

// デフォルトモデル
func fDefault(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}

// たまごのケース
func fEggCase(x, y float64) float64 {
	sx := math.Sin(x * 2 * math.Pi / xyrange * 2)
	sy := math.Sin(y * 2 * math.Pi / xyrange * 2)
	r := math.Hypot(sx, sy)
	return r / 3
}

// モーグル
func fMoguls(x, y float64) float64 {
	var sy float64
	sx := math.Sin(y*2*math.Pi/cells*6)*math.Sin(x*2*math.Pi/cells*6) + math.Sin(math.Pi/2)
	r := math.Hypot(sx, sy)
	return r / 4
}

// 乗馬用の鞍
func fSaddle(x, y float64) float64 {
	var sx, sy float64
	sx = math.Cos(x*2*math.Pi/xyrange) + math.Sin(x*2*math.Pi/xyrange*2-math.Pi/3) + math.Sin(math.Pi/2)
	if sx < 0 {
		sx = 0
	}
	filter := math.Cos(y*2*math.Pi/xyrange) + 0.3*math.Sin(y*2*math.Pi/xyrange*2-math.Pi/3) + math.Sin(math.Pi/2)
	if filter < 0 {
		filter = 0
	}
	sx = sx * 0.5 * filter
	r := math.Hypot(sx, sy)
	return r / 3
}
