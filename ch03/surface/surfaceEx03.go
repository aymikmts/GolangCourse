package surface

import (
	"fmt"
)

// 各モデルの高さの最小値と最大値を定義
const (
	defaultMin = -0.2173
	defaultMax = 0.9851
	eggCaseMin = 0
	eggCaseMax = 0.4705
	mogulsMin  = 0
	mogulsMax  = 0.5000
	saddleMin  = 0
	saddleMax  = 0.7386
)

var zMin, zMax float64 = 100, -100

func calcColor(z float64) string {
	// RGBのR, Bに対し、値を算出
	r := z*(-0xff)/(zMin-zMax) + (0xff*zMin)/(zMin-zMax)
	b := z*(0xff)/(zMin-zMax) - (0xff*zMax)/(zMin-zMax)
	ret := fmt.Sprintf("#%02x00%02x", int(r), int(b))
	return ret
}
