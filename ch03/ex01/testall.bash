#!/bin/bash
cd `dirname $0`

cd ../surface
go test -run=TestCorner
