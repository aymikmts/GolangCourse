#!/bin/bash
cd `dirname $0`

cd ../comma
go test -v -run="TestCommaWithBuffer"
