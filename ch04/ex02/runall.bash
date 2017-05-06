#!/bin/bash
cd `dirname $0`

echo --- sha256 ---
go run main.go x
echo --- sha384 ---
go run main.go -type sha384 x
echo --- sha512 ---
go run main.go -type sha512 x

