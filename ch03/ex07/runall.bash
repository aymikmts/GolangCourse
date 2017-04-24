#!/bin/bash
cd `dirname $0`

go run main.go -fractal newton > out_newton_monochrome.png
go run main.go -fractal newton -color > out_newton_color.png
