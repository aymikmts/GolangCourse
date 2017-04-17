// mainは3-D面の関数のSVGレンダリングを計算します。
package main

import (
	"GolangCourse/ch03/surface"
	"os"
)

func main() {
	surface.PrintXML(os.Stdout)
}
