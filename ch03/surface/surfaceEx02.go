package surface

import (
	"math"
)

// ModelTypeによって描画するモデルをswitch
func f(x, y float64) (r float64) {
	switch Model {
	case DEFAULT:
		r = fDefault(x, y)

	case EGGCASE:
		r = fEggCase(x, y)

	case MOGULS:
		r = fMoguls(x, y)

	case SADDLE:
		r = fSaddle(x, y)

	default:
		r = fDefault(x, y)
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
	sx := math.Sin(x * 2 * math.Pi / xyrange * 2)
	sy := math.Sin(y * 2 * math.Pi / xyrange * 2)
	r := math.Hypot(sx, sy)
	return r / 3
}

// モーグル
func fMoguls(x, y float64) float64 {
	var sy float64
	sx := math.Sin(y*2*math.Pi/cells*6)*math.Sin(x*2*math.Pi/cells*6) + math.Sin(math.Pi/2)
	r := math.Hypot(sx, sy)
	return r / 4
}

// 乗馬用の鞍
func fSaddle(x, y float64) float64 {
	var sx, sy float64
	sx = math.Cos(x*2*math.Pi/xyrange) + math.Sin(x*2*math.Pi/xyrange*2-math.Pi/3) + math.Sin(math.Pi/2)
	if sx < 0 {
		sx = 0
	}
	filter := math.Cos(y*2*math.Pi/xyrange) + 0.3*math.Sin(y*2*math.Pi/xyrange*2-math.Pi/3) + math.Sin(math.Pi/2)
	if filter < 0 {
		filter = 0
	}
	sx = sx * 0.5 * filter
	r := math.Hypot(sx, sy)
	return r / 3
}
