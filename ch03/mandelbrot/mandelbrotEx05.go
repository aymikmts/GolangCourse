package mandelbrot

import (
	"image/color"
	"math/cmplx"
)

// フルカラーのマンデルブロ集合の生成
func mandelbrotColor(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			blue := uint8(127 - contrast*n)
			red := uint8(255 - contrast*n)
			return color.YCbCr{200, blue, red}
		}
	}
	return color.Black
}
