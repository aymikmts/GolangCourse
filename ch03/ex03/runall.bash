#!/bin/bash
cd `dirname $0`

go run main.go -color > out_Color.xml
go run main.go -color -model EggCase > out_EggCase_Color.xml
go run main.go -color -model Moguls > out_Moguls_Color.xml
go run main.go -color -model Saddle > out_Saddle_Color.xml
