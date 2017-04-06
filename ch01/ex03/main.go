// mainは、echoの非効率バージョン(echo2)とstringパッケージを使ったバージョンを実行します。
package main

import (
	"fmt"
	"time"
	"strconv"

	"./echo2"
	"./echo3"
)

func main() {
	var args []string
	for i:=0; i<10; i++ {
		s := strconv.Itoa(i)
		args = append(args, s)
	}

	start1 := time.Now()
	echo2.Echo(args)
	ms1 := time.Since(start1).Nanoseconds()
	fmt.Printf("echo2 time: %v nanoseconds\n\n", ms1)

	start2 := time.Now()
	echo3.Echo(args)
	ms2 := time.Since(start2).Nanoseconds()
	fmt.Printf("echo3 time: %v nanoseconds\n", ms2)

}
