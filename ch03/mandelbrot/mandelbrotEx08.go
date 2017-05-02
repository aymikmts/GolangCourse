package mandelbrot

import (
	"image/color"
	"math/big"
	"math/cmplx"
)

// フラクタルの精度の種類を定義
type FormatType int

const (
	DEFAULT FormatType = iota
	CMPLX64
	CMPLX128
	BIGFLOAT
	BIGRAT
)

var FormatCompTest FormatType = DEFAULT
var IterationsVal = 200

func compCmplxFormat(px, py int) color.Color {
	y := float64(py)/Height*(YMax-YMin) + YMin
	x := float64(px)/Width*(XMax-XMin) + XMin

	var color color.Color

	switch FormatCompTest {
	case CMPLX64:
		color = formatCmplx64(x, y)
	case CMPLX128:
		color = formatCmplx128(x, y)
	case BIGFLOAT:
		color = formatBigFloat(x, y)
	case BIGRAT:
		color = formatBigRat(x, y)
	default:
		z := complex(x, y)
		color = mandelbrot(z)
	}

	return color
}

// complex64を使用してマンデルブロ集合を生成する。
func formatCmplx64(x, y float64) color.Color {
	z := complex64(complex(x, y))
	var iterations = uint8(IterationsVal)
	const contrast = 15

	var v complex64
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(complex128(v)) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

// complex128を使用してマンデルブロ集合を生成する(デフォルト)。
func formatCmplx128(x, y float64) color.Color {
	z := complex(x, y)
	var iterations = uint8(IterationsVal)
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

// big.Floatを使用してマンデルブロ集合を生成する。
// bigではComplexが扱えないため、実数部・虚数部に分けて算出する。
func formatBigFloat(x, y float64) color.Color {
	var iterations = uint8(IterationsVal)
	const contrast = 15

	var realZ, imagZ big.Float
	realZ.SetFloat64(x)
	imagZ.SetFloat64(y)
	var a, b big.Float
	for n := uint8(0); n < iterations; n++ {
		// 実数部計算
		// a_n+1 = a_n^2 - b_n^2 + realZ
		var a1, a2, a3, aa big.Float
		aa.Add(a3.Sub(a1.Mul(&a, &a), a2.Mul(&b, &b)), &realZ)

		// 虚数部計算
		// b_n+1 = 2 * a_n * b_n + imagZ
		var b1, b2, bb, c2 big.Float
		c2.SetFloat64(2.0)
		bb.Add(b2.Mul(b1.Mul(&c2, &a), &b), &imagZ)

		a = aa
		b = bb

		// a^2 + b^2 > 4 かどうか
		var def, aa2, bb2, comp big.Float
		def.Add(aa2.Mul(&a, &a), bb2.Mul(&b, &b))
		comp.SetFloat64(4.0)
		if def.Cmp(&comp) > 0 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}

// big.Ratを使用してマンデルブロ集合を生成する。
// bigではComplexが扱えないため、実数部・虚数部に分けて算出する。
func formatBigRat(x, y float64) color.Color {
	var iterations = uint8(IterationsVal)
	const contrast = 15

	var realZ, imagZ big.Rat
	realZ.SetFloat64(x)
	imagZ.SetFloat64(y)
	var a, b big.Rat

	for n := uint8(0); n < iterations; n++ {
		// 実数部計算
		// a_n+1 = a_n^2 - b_n^2 + realZ
		var a1, a2, a3, aa big.Rat
		aa.Add(a3.Sub(a1.Mul(&a, &a), a2.Mul(&b, &b)), &realZ)

		// 虚数部計算
		// b_n+1 = 2 * a_n * b_n + imagZ
		var b1, b2, bb, c2 big.Rat
		c2.SetFloat64(2.0)
		bb.Add(b2.Mul(b1.Mul(&c2, &a), &b), &imagZ)

		a = aa
		b = bb

		// a^2 + b^2 > 4 かどうか
		var def, aa2, bb2, comp big.Rat
		def.Add(aa2.Mul(&a, &a), bb2.Mul(&b, &b))
		comp.SetFloat64(4.0)
		if def.Cmp(&comp) > 0 {
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
}
