#!/bin/bash
cd `dirname $0`

go run main.go 0 1 2 3 4 5
go run main.go 10 20 30
go run main.go test1 test2 test3 test4 test5
