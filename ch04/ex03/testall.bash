#!/bin/bash
cd `dirname $0`

cd ../rev
go test -v -run="TestReverseByPointer"
