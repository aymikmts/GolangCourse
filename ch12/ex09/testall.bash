#!/bin/bash
cd `dirname $0`

# ex09: decode.go Token()
cd sexpr
go test -v -run="TestToken"