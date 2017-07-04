#!/bin/bash
cd `dirname $0`

cd bank
go test -v -run="TestWithdraw"
