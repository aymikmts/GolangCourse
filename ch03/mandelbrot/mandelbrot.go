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

// アンチエイリアスをかけるか否かを選択
var IsAntiAlias bool = false

// 描画するフラクタルの種類を定義
type FractalType int

const (
	MANDELBROT FractalType = iota // マンデルブロ集合
	NEWTON                        // ニュートン法によるフラクタル
)

var Fractal FractalType = MANDELBROT

// 定義値
const (
	XMin, YMin, XMax, YMax = -2, -2, +2, +2 // 軸範囲
	Width, Height          = 1024, 1024     // 描画サイズ
)

// フラクタル画像の描画
func DrawFractal(out io.Writer) {
	img := image.NewRGBA(image.Rect(0, 0, Width, Height))
	for py := 0; py < Height; py++ {
		for px := 0; px < Width; px++ {
			// 座標値に色情報をセット
			if FormatCompTest == DEFAULT {
				// default
				img.Set(px, py, calcPixelColor(px, py))
			} else {
				// ex08で使用
				img.Set(px, py, compCmplxFormat(px, py))
			}
		}
	}

	// ex06で使用
	if IsAntiAlias {
		antiAlias(img)
	}
	png.Encode(out, img) // 注意: エラーを無視
}

func calcPixelColor(px, py int) color.Color {
	y := float64(py)/Height*(YMax-YMin) + YMin
	x := float64(px)/Width*(XMax-XMin) + XMin
	z := complex(x, y)

	var color color.Color
	switch Fractal {
	case MANDELBROT:
		color = mandelbrot(z) // 色付き版の実装は"mandelbrotEx05.go"
	case NEWTON:
		color = newton(z) // 実装は"mandelbrotEx07.go"
	default:
		color = mandelbrot(z)
	}
	return color
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
