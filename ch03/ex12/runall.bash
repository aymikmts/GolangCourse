#!/bin/bash
cd `dirname $0`

go run main.go "hello, world!" "horlo, !welld"
go run main.go "Hello, World!" "horlo, !welld"
go run main.go "こんにちは、世界" "こ世ん界に、ちは"
go run main.go test foobar
