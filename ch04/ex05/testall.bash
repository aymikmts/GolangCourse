#!/bin/bash
cd `dirname $0`

cd ../dedup
go test -v -run="TestDedup"
