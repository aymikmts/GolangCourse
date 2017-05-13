#!/bin/bash
cd `dirname $0`

cd ../charcount
go test -v -run="TestCharType"
