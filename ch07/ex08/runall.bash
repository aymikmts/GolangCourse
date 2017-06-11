#!/bin/bash
cd `dirname $0`

echo --- no sorting case --
go run main.go

echo --- sorting by Track ---
go run main.go Track

echo --- sorting by Track and Year ---
go run main.go Track Year
