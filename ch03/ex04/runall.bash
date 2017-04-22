#!/bin/bash
cd `dirname $0`

go run main.go &
open "http://localhost:8000/?color=gradient&modelType=Moguls&width=640&height=480&cells=50&xyrange=40&zscale=100"
