// popcountは、xにセットされているビット数を返します。
package popcountroop

// pc[i]はiのポピュレーションカウントです。
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCountはxのポピュレーションカウント(1が設定されているビット数)を返します
// カウントはループによって行います。
func PopCount(x uint64) int {
	var ret int
	var i uint
	for i = 0; i < 8; i++ {
		ret += int(pc[byte(x>>(i*8))])
	}
	return ret
}
