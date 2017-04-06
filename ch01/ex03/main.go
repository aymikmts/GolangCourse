// mainは、echoの非効率バージョン(echo2)とstringパッケージを使ったバージョンを実行します。
package main

import (
	"fmt"
	"os"
	"time"

	"./echo2"
	"./echo3"
)

func main() {
	start := time.Now()
	echo2.Echo(os.Args[1:])
	ms := time.Since(start).Nanoseconds()
	fmt.Printf("echo2 time: %vs\n\n", ms)

	start = time.Now()
	echo3.Echo(os.Args[1:])
	ms = time.Since(start).Nanoseconds()
	fmt.Printf("echo3 time: %vs\n", ms)

}
