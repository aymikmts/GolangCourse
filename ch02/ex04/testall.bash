#!/bin/bash
cd `dirname $0`

echo --- TEST PopCountBitShift ---
cd popcountbitshift
go test -bench=.

echo --- TEST PopCount \(with Table\) ---
cd ../../ex03/popcount
go test -bench=.

