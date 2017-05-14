#!/bin/bash
cd `dirname $0`

go run main.go repo:golang/go is:open json encoder &
open -a "/Applications/Firefox.app" http://localhost:8000

