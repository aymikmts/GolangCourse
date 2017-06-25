#!/bin/bash
cd `dirname $0`

go build -o netcat3 netcat3.go
go build -o reverb1 reverb1.go

./reverb1 &
./netcat3
