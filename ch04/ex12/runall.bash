#!/bin/bash
cd `dirname $0`

echo DOWNLOAD 100
go run main.go -download 100

echo -e "\nSHOW INDEX"
go run main.go -show

echo -e "\nSEARCH \"Family\""
go run main.go -search Family

