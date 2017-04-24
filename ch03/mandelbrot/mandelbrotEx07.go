package mandelbrot

import (
	"image/color"
	"math/cmplx"
)

func newton(z complex128) color.Color {
	if IsColoring {
		return newtonColor(z)
	} else {
		return newtonMono(z)
	}
}

// モノクロのニュートンフラクタルの生成
func newtonMono(z complex128) color.Color {
	const iterations = 37
	const contrast = 7
	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(z*z*z*z-1) < 1e-6 {
			return color.Gray{255 - contrast*i}
		}
	}
	return color.Black
}

// フルカラーのニュートンフラクタルの生成
func newtonColor(z complex128) color.Color {
	const iterations = 37
	const contrast = 7
	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(z*z*z*z-1) < 1e-6 {
			blue := uint8(127 - contrast*i)
			red := uint8(255 - contrast*i)
			return color.YCbCr{200, blue, red}
		}
	}
	return color.Black
}
