#!/bin/bash
cd `dirname $0`

echo --- TEST PopCount \(with Table\) ---
cd popcount
go test -bench=.

echo --- TEST PopRoop ---
cd ../popcountroop
go test -bench=.

