#!/bin/bash
cd `dirname $0`

go run main.go > out_noantialias.png
go run main.go -antialias > out_antialias.png

go run main.go -color > out_noantialias_color.png
go run main.go -antialias -color > out_antialias_color.png
