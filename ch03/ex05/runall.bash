#!/bin/bash
cd `dirname $0`

go run main.go > out_monochrome.png
go run main.go -color > out_color.png