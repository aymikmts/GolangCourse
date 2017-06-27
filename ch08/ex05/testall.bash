#!/bin/bash/
cd `dirname $0`

cd mandelbrot
go test -bench=.

cd ../surface
go test -bench=.
