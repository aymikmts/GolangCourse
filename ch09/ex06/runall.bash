#!/bin/bash
cd `dirname $0`

echo build program
echo このプログラムはフィボナッチ数列を並列計算しています。
go build -o ex06 main.go

echo -------
echo GOMAXPROCS=1
export GOMAXPROCS=1
time ./ex06

echo
echo -------
echo GOMAXPROCS=2
export GOMAXPROCS=2
time ./ex06

echo
echo -------
echo GOMAXPROCS=3
export GOMAXPROCS=3
time ./ex06

echo
echo -------
echo GOMAXPROCS=4
export GOMAXPROCS=4
time ./ex06


echo
echo -------
echo GOMAXPROCS=8
export GOMAXPROCS=8
time ./ex06


