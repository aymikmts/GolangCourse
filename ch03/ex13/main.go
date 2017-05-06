// Ex13はKB, MB, …, YBまでのconst宣言を表示します。
package main

import (
	"fmt"
	"strconv"

	"../comma"
	//"GolangCourse/ch03/comma"
)

const (
	_   = 1 << (10 * iota)
	KiB // 1024
	MiB // 1048576
	GiB // 1073741824
	TiB // 1099511627776
	PiB // 1125899906842624
	EiB // 1152921504606846976
	ZiB //
	YiB //
)

const (
	KB = 1000
	MB = KB * KB
	GB = MB * KB
	TB = GB * KB
	PB = TB * KB
	EB = PB * KB
	ZB = EB * KB
	YB = ZB * KB
)

func main() {
	fmt.Printf("KB: %s\n", comma.CommaWithBuffer(strconv.Itoa(KB)))
	fmt.Printf("MB: %s\n", comma.CommaWithBuffer(strconv.Itoa(MB)))
	fmt.Printf("GB: %s\n", comma.CommaWithBuffer(strconv.Itoa(GB)))
	fmt.Printf("TB: %s\n", comma.CommaWithBuffer(strconv.Itoa(TB)))
	fmt.Printf("PB: %s\n", comma.CommaWithBuffer(strconv.Itoa(PB)))
	fmt.Printf("EB: %s\n", comma.CommaWithBuffer(strconv.Itoa(EB)))
	//fmt.Printf("ZB: %s\n", comma.CommaWithBuffer(strconv.Itoa(ZB)))
	//fmt.Printf("YB: %s\n", comma.CommaWithBuffer(strconv.Itoa(YB)))
}
