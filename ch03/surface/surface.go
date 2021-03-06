// Surfaceは3-D面の関数のSVGレンダリングを計算します。
package surface

import (
	"fmt"
	"io"
	"math"
	"os"
)

// 描画するものの種類を定義
type ModelType int

const (
	DEFAULT ModelType = iota
	EGGCASE
	MOGULS
	SADDLE
)

// 描画モデルの種類
var Model ModelType = DEFAULT // デフォルトモデル

// 着色有無
var IsGradientColor bool = false // グラデーションの有無
var FillColor = "white"          // デフォルトの色

const (
	width, height = 600, 320            // キャンバスの大きさ(画素数)
	cells         = 100                 // 格子のます目の数
	xyrange       = 30.0                // 軸の範囲 (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // x単位 および y単位あたりの画素数
	zscale        = height * 0.4        // z単位あたりの画素数
	angle         = math.Pi / 6         // x, y軸の角度 (=30度)
)

var (
	Width, Height float64 = width, height
	Cells                 = cells
	XYRange               = xyrange
	ZScale                = zscale
)
var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30度), cos(30度)

func PrintXML(out io.Writer) {
	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: gray; stroke-width: 0.7' "+
		"width='%d' height='%d'>", Width, Height)
	for i := 0; i < Cells; i++ {
		for j := 0; j < Cells; j++ {
			// [ex01]戻り値okがfalseだったときは、スキップする。
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

			// [ex03]fillパラメータを追加。
			// corner()内のcalcColor()で算出したカラーコードを記述する
			fmt.Fprintf(out, "<polygon points='%g,%g %g,%g %g,%g %g,%g' fill=\"%s\"/>\n",
				ax, ay, bx, by, cx, cy, dx, dy, FillColor)
		}
	}
	fmt.Fprintln(out, "</svg>")
}

func corner(i, j int) (sx float64, sy float64, ok bool) {
	// ます目(i,j)のかどの点(x, y)を見つける。
	x := XYRange * (float64(i)/float64(Cells) - 0.5)
	y := XYRange * (float64(j)/float64(Cells) - 0.5)

	// 面の高さzを計算する。
	z := f(x, y)
	if math.IsNaN(z) {
		// [ex01]値がNaNだった場合はStderr出力し、falseを返す。
		fmt.Fprintf(os.Stderr, "[%d, %d]z := f(%.5f, %.5f) : Value is NaN!\n", i, j, x, y)
		ok = false
		return
	}

	// [ex03]zに応じて色を算出する
	if IsGradientColor {
		FillColor = calcColor(z)
	}

	// (x, y, z)を2-D SVGキャンバス(sx, sy)へ等角的に投影。
	sx = Width/2 + (x-y)*cos30*xyscale
	sy = Height/2 + (x+y)*sin30*xyscale - z*ZScale
	ok = true
	return
}
