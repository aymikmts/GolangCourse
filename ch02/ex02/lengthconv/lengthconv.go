// lengthconvパッケージはフィートとメートルの長さ変換を行います。
package lengthconv

import "fmt"

type Feet float64
type Meter float64

const (
	convFToMVal float64 = 0.3048
)

func (f Feet) String() string  { return fmt.Sprintf("%gfeet", f) }
func (m Meter) String() string { return fmt.Sprintf("%gmeter", m) }
