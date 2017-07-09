#!/bin/bash
cd `dirname $0`

echo build programs.
go build -o fetch fetch.go
go build -o xmlselect xmlselect.go

echo Run
./fetch http://www.w3.org/TR/2006/REC-xml11-20060816 | ./xmlselect div div class="div1" h2

