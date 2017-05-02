package mandelbrot

import (
	"fmt"
	"image"
	"image/png"
	"io"
	"math"
	"net/http"
	"strconv"
)

// 指定パラメータ
const (
	keyX       = "x"
	keyY       = "y"
	keyScale   = "scale"
	keyColor   = "color"
	keyFractal = "fractal"
)

var XVal float64 = 0.0
var YVal float64 = 0.0
var Scale float64 = 1.0

func Handler(w http.ResponseWriter, r *http.Request) {
	// パラメータを初期化
	XVal = 0.0
	YVal = 0.0
	Scale = 1.0

	// パラメータを読み取る
	for key, values := range r.URL.Query() {
		switch key {
		case keyX:
			setXVal(w, values[0])
		case keyY:
			setYVal(w, values[0])
		case keyScale:
			setScale(w, values[0])
		case keyColor:
			setColor(w, values[0])
		case keyFractal:
			setFractal(w, values[0])
		}
	}

	// 画像を描画
	DrawFractalHttp(w)
}

// x座標の設定
func setXVal(w http.ResponseWriter, val string) {
	param, err := strconv.ParseFloat(val, 64)
	if err != nil {
		fmt.Fprintf(w, "invalid value: %v\n", val)
	}
	XVal = float64(param)
}

// y座標の設定
func setYVal(w http.ResponseWriter, val string) {
	param, err := strconv.ParseFloat(val, 64)
	if err != nil {
		fmt.Fprintf(w, "invalid value: %v\n", val)
	}
	YVal = float64(param)
}

// 拡大値の設定
func setScale(w http.ResponseWriter, val string) {
	param, err := strconv.ParseFloat(val, 64)
	if err != nil {
		fmt.Fprintf(w, "invalid value: %v\n", val)
	}

	Scale = math.Sqrt(float64(param))
	xMin = XMin / Scale
	xMax = XMax / Scale
	yMin = YMin / Scale
	yMax = YMax / Scale
}

// 色付きか否かを設定
func setColor(w http.ResponseWriter, val string) {
	if val == "true" {
		IsColoring = true
	} else {
		IsColoring = false
	}
}

// 描画するフラクタルの種類を設定
func setFractal(w http.ResponseWriter, val string) {
	if val == "newton" {
		Fractal = NEWTON
	} else {
		Fractal = MANDELBROT
	}
}

// フラクタル画像の描画
func DrawFractalHttp(out io.Writer) {

	// 画像の描画範囲を広げる。
	width = int(Width * Scale)
	height = int(Height * Scale)

	tmpImg := image.NewRGBA(image.Rect(-Width/2, -Height/2, width+Width/2, height+Height/2))
	for py := 0; py < height; py++ {
		for px := 0; px < width; px++ {
			// 座標値に色情報をセット
			tmpImg.Set(px, py, calcPixelColor(px, py))
		}
	}

	// 拡大画像から表示部分を切り取る
	xPixel := int(Width/4.0*Scale*XVal + Width/2.0*Scale - 1.0/2.0*Width)
	yPixel := int(Height/4.0*Scale*YVal + Height/2.0*Scale - 1.0/2.0*Height)
	subImg := tmpImg.SubImage(image.Rect(xPixel, yPixel, xPixel+Width, yPixel+Height))

	png.Encode(out, subImg) // 注意: エラーを無視
}
