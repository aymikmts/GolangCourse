#!/bin/bash
cd `dirname $0`

echo !!! TODO:大文字小文字の区別をしない、ピリオドなどの記号は取り除く !!!
go run main.go < input
