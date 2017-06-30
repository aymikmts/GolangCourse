#!/bin/bash
cd `dirname $0`

go build -o reverb reverb.go
go build -o netcat netcat.go
echo FINISH BUILD.
echo PLEASE INPUT WORDS.

./reverb &
./netcat
