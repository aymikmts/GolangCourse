#!/bin/bash/
cd `dirname $0`
cd popcount

go test -bench=.

# Go 1.8, 1.3GHz Intel Core i5
# BenchmarkPopCount_0-4                   2000000000               0.41 ns/op
# BenchmarkPopCountByShifiting_0-4        20000000                98.9 ns/op
# BenchmarkPopCountByClearing_0-4         1000000000               2.49 ns/op
#
# BenchmarkPopCount_16c-4                 2000000000               0.45 ns/op
# BenchmarkPopCountByShifiting_16c-4      20000000               101 ns/op
# BenchmarkPopCountByClearing_16c-4       100000000               15.0 ns/op
#
# BenchmarkPopCount_32c-4                 2000000000               0.42 ns/op
# BenchmarkPopCountByShifiting_32c-4      20000000               106 ns/op
# BenchmarkPopCountByClearing_32c-4       50000000                29.4 ns/op
#
# BenchmarkPopCount_64c-4                 2000000000               0.41 ns/op
# BenchmarkPopCountByShifiting_64c-4      20000000                96.1 ns/op
# BenchmarkPopCountByClearing_64c-4       20000000                69.3 ns/op