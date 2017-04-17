#!/bin/bash
cd `dirname $0`

go run main.go -model EggCase > out_EggCase.xml
go run main.go -model Moguls > out_Moguls.xml
go run main.go -model Saddle > out_Saddle.xml
