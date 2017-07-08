#!/bin/bash
cd `dirname $0`

echo depth=0
go run main.go -depth=0 http://gopl.io

echo 
echo depth=1
go run main.go -depth=1 http://gopl.io

