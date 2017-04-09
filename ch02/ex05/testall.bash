#!/bin/bash

CURRENT_DIR=$PWD
cd `dirname $0`

echo --- TEST PopCount \(with Table\) ---
cd $CURRENT_DIR
cd ../ex03/popcount
go test -bench=.

echo --- TEST PopRoop ---
cd $CURRENT_DIR
cd ../ex03/popcountroop
go test -bench=.

echo --- TEST PopCountBitShift ---
cd $CURRENT_DIR
cd ../ex04/popcountbitshift
go test -bench=.

echo --- TEST PopCountAndOperation ---
cd $CURRENT_DIR
cd popcountandoperation
go test -bench=.

