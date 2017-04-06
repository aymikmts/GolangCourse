// mainは、echoの非効率バージョン(echo2)とstringパッケージを使ったバージョンを実行します。
package main

import (
	"os"

	"./echo2"
	"./echo3"
)

func main() {
	echo2.Echo(os.Args[1:])
	echo3.Echo(os.Args[1:])
}
