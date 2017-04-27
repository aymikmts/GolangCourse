#!/bin/bash
# TODO: 拡大率変えてみる、メモリ量の算出
cd `dirname $0`

go run main.go -format complex64 > out_cmplx64.png
go run main.go -format complex128 > out_cmplx128.png
go run main.go -format big.Float > out_bigFloat.png
# 時間がかかりすぎるためコメントアウト
#go run main.go -format big.Rat > out_bigRat.png