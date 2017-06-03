#!/bin/bash
cd `dirname $0`

echo !! SHOW DIRECTORIES !!
go run main.go -d ../ex14

echo !! SHOW TOPOSORT DEPENDENCY !!
go run main.go -t algorithms

echo !! SHOW CRAWL URL !!
go run main.go -u http://golang.org
