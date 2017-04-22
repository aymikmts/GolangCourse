#!/bin/bash
cd `dirname $0`

go run main.go -gradient > out_Color.xml
go run main.go -gradient -model EggCase > out_EggCase_Color.xml
go run main.go -gradient -model Moguls > out_Moguls_Color.xml
go run main.go -gradient -model Saddle > out_Saddle_Color.xml
