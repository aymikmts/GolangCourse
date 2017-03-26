// weightconvパッケージはポンド(Pound)とキログラム(KiloGram)の重さ変換を行います
package weightconv

import "fmt"

type Pound float64
type KiloGram float64

const (
	convPToKgVal float64 = 0.453592
)

func (p Pound) String() string     { return fmt.Sprintf("%gpound", p) }
func (kg KiloGram) String() string { return fmt.Sprintf("%gkg", kg) }
