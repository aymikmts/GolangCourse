package surface

import (
	"math"
)

// [ex02]引数で与えられた値によって描画するモデルをswitch
// [ex03]各Model高さの最小値、最大値をxMin, xMaxにセットする
func f(x, y float64) (r float64) {
	switch Model {
	case EGGCASE:
		r = fEggCase(x, y)
		zMin = eggCaseMin
		zMax = eggCaseMax

	case MOGULS:
		r = fMoguls(x, y)
		zMin = mogulsMin
		zMax = mogulsMax

	case SADDLE:
		r = fSaddle(x, y)
		zMin = saddleMin
		zMax = saddleMax

	default:
		r = fDefault(x, y)
		zMin = defaultMin
		zMax = defaultMax
	}
	return
}

// デフォルトモデル
func fDefault(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}

// たまごのケース
func fEggCase(x, y float64) float64 {
	sx := math.Sin(x * 2 * math.Pi / XYRange * 2)
	sy := math.Sin(y * 2 * math.Pi / XYRange * 2)
	r := math.Hypot(sx, sy)
	return r / 3
}

// モーグル
func fMoguls(x, y float64) float64 {
	var sy float64
	sx := math.Sin(y*2*math.Pi/float64(Cells)*6)*math.Sin(x*2*math.Pi/float64(Cells)*6) + math.Sin(math.Pi/2)
	r := math.Hypot(sx, sy)
	return r / 4
}

// 乗馬用の鞍
func fSaddle(x, y float64) float64 {
	var sx, sy float64
	sx = math.Cos(x*2*math.Pi/XYRange) + math.Sin(x*2*math.Pi/XYRange*2-math.Pi/3) + math.Sin(math.Pi/2)
	if sx < 0 {
		sx = 0
	}
	filter := math.Cos(y*2*math.Pi/XYRange) + 0.3*math.Sin(y*2*math.Pi/XYRange*2-math.Pi/3) + math.Sin(math.Pi/2)
	if filter < 0 {
		filter = 0
	}
	sx = sx * 0.5 * filter
	r := math.Hypot(sx, sy)
	return r / 3
}
