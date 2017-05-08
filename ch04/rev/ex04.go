package rev

// Rotateはint型のスライスをrot分回転させます。
// rotが正の値のときは左に、負の値は右に回転します。
func Rotate(s []int, rot int) {
	if rot < 0 {
		rot = len(s) + rot
	}
	// Rotate分をコピー
	tmp := make([]int, rot)
	copy(tmp, s[:rot])

	for i := 0; i < len(s); i++ {
		if i < len(s)-rot {
			s[i] = s[i+rot]
		} else {
			s[i] = tmp[i-len(s)+rot]
		}
	}
}
