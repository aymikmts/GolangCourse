#!/bin/bash
cd `dirname $0`

go run main.go 2500000

# macOS Sierra ver10.12.6
# CPU: 1.3 GHz Intel Core i5
# Memory: 8GB
# 使用済みメモリ: 3.72GB

# go run main.go 2500000
# Finish. n:2500000
# [2500000 goroutines]
# total: 25.866710389s average: 10346 ns

# go run main.go 3000000
# Finish. n:3000000
# [3000000 goroutines]
# total: 27.841889586s average: 9280 ns

# go run main.go 5000000
# Finish. n:5000000
# [5000000 goroutines]
# total: 1m10.942613193s average: 14188 ns

# go run main.go 10000000
# Finish. n:10000000
# [10000000 goroutines]
# total: 2m51.408967897s average: 17140 ns

# 2,420,000辺りでスワップが起き始めていた