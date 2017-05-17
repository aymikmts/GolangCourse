#!/bin/bash
cd `dirname $0`

# 正常ケース
go run main.go harry potter

# エラーケース(映画名がヒットしなかったとき)
go run main.go hurry potter

