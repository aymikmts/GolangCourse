package main

import (
	"GolangCourse/ch03/mandelbrot"
	"io"
	"os"
)

var output io.Writer = os.Stdout

func main() {
	mandelbrot.DrawMandelbrot(output)
}
