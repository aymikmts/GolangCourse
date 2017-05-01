package mandelbrot

import (
	"image"
	"image/color"
)

// 4つのサブピクセルに分割してカラー値の平均を算出する手法のスーパーサンプリングを実行します。
func antiAlias(input *image.RGBA) *image.RGBA {
	out := image.NewRGBA(image.Rect(0, 0, Width, Height))

	// 高画質化した画像の2pixel四方分の色平均を求め、
	// 出力画像の1画素の値とする。
	for j := 0; j < Height-1; j++ {
		for i := 0; i < Width-1; i++ {
			// 2pixel四方の色を取得
			r1, g1, b1, _ := input.At(i, j).RGBA()
			r2, g2, b2, _ := input.At(i+1, j).RGBA()
			r3, g3, b3, _ := input.At(i, j+1).RGBA()
			r4, g4, b4, _ := input.At(i+1, j+1).RGBA()

			rTmp := (r1 + r2 + r3 + r4) / 4
			// 32bitから8bitに変換
			r := uint8(rTmp * 0xff / 0xffff)

			gTmp := (g1 + g2 + g3 + g4) / 4
			// 32bitから8bitに変換
			g := uint8(gTmp * 0xff / 0xffff)

			bTmp := (b1 + b2 + b3 + b4) / 4
			// 32bitから8bitに変換
			b := uint8(bTmp * 0xff / 0xffff)

			color := color.RGBA{r, g, b, 0xff}
			out.SetRGBA(i, j, color)
		}
	}
	return out
}
