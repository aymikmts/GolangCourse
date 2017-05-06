#!/bin/bash
cd `dirname $0`

echo --- [test][test] ---
go run main.go test test
echo --- [x][X] ---
go run main.go x X

