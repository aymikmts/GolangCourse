#!/bin/bash
cd `dirname $0`

cd intset
go test -bench=.

# Go1.8, 1.3GHz Intel Core i5
# Benchmark_IntSet_Add_Word8-4                      300000              5527 ns/op
# Benchmark_IntSet_Add_Word16-4                     200000              5883 ns/op
# Benchmark_IntSet_Add_Word32-4                      50000             25723 ns/op
#
# Benchmark_IntSet_UnionWith_Word8-4                100000             11356 ns/op
# Benchmark_IntSet_UnionWith_Word16-4               100000             12229 ns/op
# Benchmark_IntSet_UnionWith_Word32-4                    2         500061504 ns/op
#
# Benchmark_IntSet_IntersectWith_Word8-4            100000             11622 ns/op
# Benchmark_IntSet_IntersectWith_Word16-4           100000             12290 ns/op
# Benchmark_IntSet_IntersectWith_Word32-4               10         130411238 ns/op
#
# Benchmark_IntSet_Add-4                            200000              5759 ns/op
# Benchmark_MapIntSet_Add-4                         100000             11849 ns/op
# Benchmark_IntSet_UnionWith-4                      100000             12105 ns/op
# Benchmark_MapIntSet_UnionWith-4                     3000           2350935 ns/op
