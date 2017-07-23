#!/bin/bash/
cd `dirname $0`

go build -o mandelbrot mandelbrot.go
go build -o ex01 main.go

echo --- no flag case ---
./mandelbrot | ./ex01 > out1.png

echo
echo --- png flag case ---
./mandelbrot | ./ex01 -f png > out2.png

echo
echo --- jpeg flag case ---
./mandelbrot | ./ex01 -f jpeg > out.jpg

echo
echo --- gif flag case ---
./mandelbrot | ./ex01 -f gif > out.gif

echo
echo --- invalid flag case ---
./mandelbrot | ./ex01 -f tiff > out.tiff