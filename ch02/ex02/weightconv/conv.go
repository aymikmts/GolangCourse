package weightconv

import "math"

// pointは四捨五入する小数点桁数
const point int = 3

// PToKgはポンドをキログラムへ変換します。
func PToKg(p Pound) KiloGram { return KiloGram(round(float64(p)*convPToKgVal, point)) }

// KgToPはキログラムをポンドへ変換します。
func KgToP(kg KiloGram) Pound { return Pound(round(float64(kg)/convPToKgVal, point)) }

// roundはfを小数点以下places桁で四捨五入します。
func round(f float64, places int) float64 {
	shift := math.Pow(10, float64(places))
	return math.Floor(f*shift+.5) / shift
}
