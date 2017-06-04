#!/bin/bash
cd `dirname $0`

go run main.go a testdata/input.html
echo -----------------
go run main.go a testdata/golang.org.html
echo -----------------
go run main.go a img h1 testdata/input.html
echo -----------------
go run main.go testdata/input.html
echo -----------------
