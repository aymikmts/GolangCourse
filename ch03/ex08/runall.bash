#!/bin/bash
# 各条件の画像はbytes.Bufferに出力し、出力前後のメモリ量をStderrで出力する。
# 計算量軽減のため、Iterations(繰り返し回数)を1に設定してある。
# 性能についてはtestall.bashを実行して、実行時間を計測する
cd `dirname $0`

echo --- complex64 ver. ---
go run main.go -format complex64

echo --- COMPLEX128 ver. ---
go run main.go -format complex128

echo --- big.Float ver. ---
go run main.go -format big.Float

# 時間がかかりすぎるためコメントアウト
echo ---big.Rat ver. ---
go run main.go -format big.Rat
