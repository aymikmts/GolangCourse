package lengthconv

import "math"

// pointは四捨五入する小数点桁数
const point int = 3

// FToMはフィートをメートルへ変換します。
func FToM(f Feet) Meter { return Meter(round(float64(f)*convFToMVal, point)) }

// MToFはメートルをフィートへ変換します。
func MToF(m Meter) Feet { return Feet(round(float64(m)/convFToMVal, point)) }

// roundはfを小数点以下places桁で四捨五入します。
func round(f float64, places int) float64 {
	shift := math.Pow(10, float64(places))
	return math.Floor(f*shift+.5) / shift
}
