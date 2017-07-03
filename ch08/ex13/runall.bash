#!/bin/bash
cd `dirname $0`

go build -o chat chat.go
go build -o netcat netcat.go
echo Finish build.

echo Start programs.
./chat &
./netcat
