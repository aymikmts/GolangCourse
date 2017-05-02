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
	HiResRatio             = 2              // Ex06で使用。高画質化比率。
)

var xMin, yMin, xMax, yMax float64 = XMin, YMin, XMax, YMax
var width, height int = Width, Height

// フラクタル画像の描画
func DrawFractal(out io.Writer) {
	outImg := image.NewRGBA(image.Rect(0, 0, Width, Height))

	// Ex06で使用
	// 有効化のときは、一時的に高画質化する
	if IsAntiAlias {
		width = Width * HiResRatio
		height = Height * HiResRatio
	}

	tmpImg := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		for px := 0; px < width; px++ {
			// 座標値に色情報をセット
			if FormatCompTest == DEFAULT {
				// default
				tmpImg.Set(px, py, calcPixelColor(px, py))
			} else {
				// ex08で使用
				tmpImg.Set(px, py, compCmplxFormat(px, py))
			}
		}
	}

	// Ex06
	// 有効化のときは、ダウンサンプリングする
	if IsAntiAlias {
		outImg = antiAlias(tmpImg)
	} else {
		outImg = tmpImg
	}

	png.Encode(out, outImg) // 注意: エラーを無視
}

func calcPixelColor(px, py int) color.Color {
	y := float64(py)/float64(height)*(yMax-yMin) + yMin
	x := float64(px)/float64(width)*(xMax-xMin) + xMin
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
