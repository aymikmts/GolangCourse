#!/bin/bash
cd `dirname $0`

go run main.go
go run main.go -temp -18C
go run main.go -temp 212F
go run main.go -temp 273.15K