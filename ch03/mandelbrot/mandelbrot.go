// MandelbrotはマンデルプロフラクタルのPNG画像を生成します。
package mandelbrot

import (
	"image"
	"image/color"
	"image/png"
	"io"
	"math/cmplx"
)

// フルカラーか否かを選択
var IsColoring bool = false

// 描画するものの種類を定義
type FractalType int

const (
	MANDELBROT FractalType = iota
	NEWTON
)

var Fractal FractalType = MANDELBROT // デフォルトはマンデルブロー集合

func DrawFractal(out io.Writer) {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// 画像の点(px, py)は複素数値zを表している。
			var color color.Color
			switch Fractal {
			case MANDELBROT:
				color = mandelbrot(z)
			case NEWTON:
				color = newton(z)
			default:
				color = mandelbrot(z)
			}
			img.Set(px, py, color)
		}
	}
	png.Encode(out, img) // 注意: エラーを無視
}

func mandelbrot(z complex128) color.Color {
	if IsColoring {
		return mandelbrotColor(z)
	} else {
		return mandelbrotMono(z)
	}
}

// モノクロのマンデルブロ集合の生成
func mandelbrotMono(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
