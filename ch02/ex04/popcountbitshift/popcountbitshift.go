// PopCountBitShiftは、引数をビットシフトしながら最下位ビットの検査を64回繰り返すことでビット数を数えます。
package popcountbitshift

// PopCountはxのポピュレーションカウント(1が設定されているビット数)を返します。
func PopCount(x uint64) int {
	var ret int
	val := x
	for i := 0; i < 64; i++ {
		ret += int(val) & 1
		val = val >> 1
	}
	return ret
}
