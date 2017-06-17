#!/bin/bash
cd `dirname $0`

# ex13は、eval/string.goに実装
cd eval
go test
