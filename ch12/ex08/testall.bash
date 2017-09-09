#!/bin/bash
cd `dirname $0`

cd sexpr
go test -v -run="TestNewDecoder"