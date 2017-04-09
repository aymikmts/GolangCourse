// PopCountAndOperationは、x&(x-1)で最下位ビットがクリアされることを利用してビット数を数えます。
package popcountandoperation

// PopCountはxのポピュレーションカウント(1が設定されているビット数)を返します。
func PopCount(x uint64) int {
	var ret int
	val := x
	for val != val&(val-1) {
		ret++
		val = val & (val - 1)
	}
	return ret
}
