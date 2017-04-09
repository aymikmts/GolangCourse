#!/bin/bash
cd `dirname $0`

echo --- TEST PopCount \(with Table\) ---
cd ../ex03/popcount
go test -bench=.

echo --- TEST PopRoop ---
cd ../popcountroop
go test -bench=.

echo --- TEST PopCountBitShift ---
cd ../ex04/popcountbitshift
go test -bench=.

