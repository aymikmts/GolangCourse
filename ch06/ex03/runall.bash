#!/bin/bash
cd `dirname $0`

echo -- UnionWith --
go run main.go -u

echo -- IntersectWith --
go run main.go -i

echo -- DifferenceWith --
go run main.go -d

echo -- SymmetricDifference --
go run main.go -s
