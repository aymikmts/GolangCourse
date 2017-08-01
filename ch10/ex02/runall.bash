#!/bin/bash
cd `dirname $0`

go build -o ex02 main.go

echo show file list of a TAR file.
./ex02 data/test.tar

echo
echo show file list of a ZIP file.
./ex02 data/test.zip

echo
echo other file.
./ex02 data/test.txt
