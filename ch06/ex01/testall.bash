#!/bin/bash
cd `dirname $0`

cd ../intset
go test -v -run="TestLen|TestRemove|TestClear|TestCopy"
