package mandelbrot

import (
	"fmt"
	"image"
	"image/color"
	"os"
)

func antiAlias(input *image.RGBA) {

	// not implemented
	fmt.Fprintf(os.Stderr, "NOT IMPLEMENTED!\n")
	os.Exit(1)

	out := image.NewRGBA(image.Rect(0, 0, Width, Height))

	for j := 0; j < Height-1; j++ {
		for i := 0; i < Width-1; i++ {
			// 4pixel四方の色を取得
			r1, g1, b1, _ := input.At(i, j).RGBA()
			r2, g2, b2, _ := input.At(i+1, j).RGBA()
			r3, g3, b3, _ := input.At(i, j+1).RGBA()
			r4, g4, b4, _ := input.At(i+1, j+1).RGBA()

			// r := uint8(float64(ulR+urR+blR+brR) / 4.0)
			// g := uint8(float64(ulG+urG+blG+brG) / 4.0)
			// b := uint8(float64(ulB+urB+blB+brB) / 4.0)
			rTmp := (r1 + r2 + r3 + r4) / 4
			if rTmp > 0xFF || rTmp < 0x00 {
				fmt.Fprintf(os.Stderr, "R is out. [%v, %v](%v, %v, %v, %v): %v\n", i, j, r1, r2, r3, r4, rTmp)
			}
			r := uint8(rTmp)

			gTmp := (g1 + g2 + g3 + g4) / 4
			if gTmp > 0xFFFF || gTmp < 0x00 {
				fmt.Fprintf(os.Stderr, "G is out. [%v, %v](%v, %v, %v, %v): %v\n", i, j, g1, g2, g3, g4, gTmp)
			}
			g := uint8(gTmp)

			bTmp := (b1 + b2 + b3 + b4) / 4
			if bTmp > 0xFFFF || bTmp < 0x00 {
				fmt.Fprintf(os.Stderr, "B is out. [%v, %v](%v, %v, %v, %v): %v\n", i, j, b1, b2, b3, b4, bTmp)
			}
			b := uint8(bTmp)

			color := color.RGBA{r, g, b, 0xff}
			out.SetRGBA(i, j, color)
		}
	}
	*input = *out
}
