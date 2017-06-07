// tempconvパッケージは摂氏(Celsius)と華氏(Fahrenheit)の温度変換を行います
package tempconv

import (
	"math"
)

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

// pointは四捨五入する小数点桁数
const point int = 3

// CToFは摂氏の温度を華氏へ変換します。
func CToF(c Celsius) Fahrenheit { return Fahrenheit(round(float64(c)*9/5+32, point)) }

// CToKは摂氏の温度を絶対温度へ変換します。
func CToK(c Celsius) Kelvin { return Kelvin(c + 273.15) }

// FToCは華氏の温度を摂氏へ変換します。
func FToC(f Fahrenheit) Celsius { return Celsius(round((float64(f)-32)*5/9, point)) }

// FToKは華氏の温度を絶対温度へ変換します。
func FToK(f Fahrenheit) Kelvin { return Kelvin(CToK(FToC(f))) }

// KToCは絶対温度を摂氏へ変換します。
func KToC(k Kelvin) Celsius { return Celsius(k - 273.15) }

// KToFは絶対温度を華氏へ変換します。
func KToF(k Kelvin) Fahrenheit { return Fahrenheit(CToF(KToC(k))) }

// roundはfを小数点以下places桁で四捨五入します。
func round(f float64, places int) float64 {
	shift := math.Pow(10, float64(places))
	return math.Floor(f*shift+.5) / shift
}
