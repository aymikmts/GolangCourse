#!/bin/bash/
cd `dirname $0`

# ex10: decode.go read(), readList()にブーリアン、浮動小数点、インテフェースを処理するように拡張
cd sexpr
go test -v -run="TestUnmarshal"