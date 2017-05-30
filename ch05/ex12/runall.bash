#!/bin/bash
cd `dirname $0`

go run main.go < testdata/test.html
#go run main.go < testdata/input.html
#go run main.go < testdata/golang.org.html
#go run main.go < testdata/theta360.com.html
