#!/bin/bash
cd `dirname $0`

go run main.go a b c ,
echo -- ERROR CASE --
go run main.go
go run main.go a

