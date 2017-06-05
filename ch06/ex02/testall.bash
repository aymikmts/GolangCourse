#!/bin/bash
cd `dirname $0`

cd ../intset
go test -v -run="TestAddAll"
