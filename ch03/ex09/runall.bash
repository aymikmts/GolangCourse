#!/bin/bash
cd `dirname $0`

go run main.go &
open "http://localhost:8000/?scale=2.0&x=0.5&y=0.5&color=true&fractal=mandelbrot"

