package digest

var pc [256]byte

// あらかじめ、8bit分のポピュレーションカウントテーブルを作成しておく
func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// CountDiffBitは2つの[32]uint8で異なるビットの数を返します。
func CountDiffBit(x1, x2 [32]uint8) int {
	var ret int
	for i, _ := range x1 {
		// 8bitごとにカウントする。
		// x1[i], x2[i]のXORをとることで、異なるビットを抽出する。
		ret += int(pc[x1[i]^x2[i]])
	}
	return ret
}
