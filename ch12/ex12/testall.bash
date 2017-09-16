#!/bin/bash
cd `dirname $0`

cd params
go test -v -run="TestCheckValue"
