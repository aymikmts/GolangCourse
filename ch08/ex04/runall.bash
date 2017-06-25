#!/bin/bash
cd `dirname $0`

go build -o netcat netcat.go
go build -o reverb reverb.go

./reverb &
./netcat
