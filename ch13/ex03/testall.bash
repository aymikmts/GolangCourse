#!/bin/bash

cd `dirname $0`

cd bzip
go test -v -run="TestBzipInParallel"
