// Ex02はSHA256/SHA384/SHA512ハッシュを表示します。
// フラグでSHA256/SHA384/SHA512を切り替えます。
// デフォルトはSHA256です。
// 実装は、"digest/ex02.go"
package main

import (
	//"GolangCourse/ch04/digest"

	"flag"
	"fmt"

	"../digest"
)

var shaTypeFlag = flag.String("type", "sha256", "sha type: sha256/sha384/sha512")

func main() {
	flag.Parse()
	switch *shaTypeFlag {
	case "sha384":
		digest.SHAFlag = digest.SHA384
	case "sha512":
		digest.SHAFlag = digest.SHA512
	default:
		digest.SHAFlag = digest.SHA256
	}

	fmt.Printf("[in]:%s\n[out]:%s\n", flag.Arg(0), digest.MakeDigest(flag.Arg(0)))

}
