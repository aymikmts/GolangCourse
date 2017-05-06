// Ex01は2つのSHA256ハッシュで異なるビットの数を数えます。
// 実装は、"digest/ex01.go"
package main

import (
	//"GolangCourse/ch04/digest"
	"crypto/sha256"
	"fmt"
	"os"

	"../digest"
)

func main() {
	x1 := sha256.Sum256([]byte(os.Args[1]))
	x2 := sha256.Sum256([]byte(os.Args[2]))
	fmt.Printf("input:\n %s\n %s\ndigest:\n %x\n %x\ncount:%d\n", os.Args[1], os.Args[2], x1, x2, digest.CountDiffBit(x1, x2))
}
