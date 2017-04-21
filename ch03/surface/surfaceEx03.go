package surface

func calcColor(z float64) string {
	//	fmt.Fprintf(os.Stderr, "z:%v\n", z)
	ret := "white"

	// 頂点が赤(#ff0000)となり谷が青(#0000ff)になるようにする。
	//	if z > 5 {
	//		ret = "#ff0000"
	//	}
	//	if z <= 0 {
	//		ret = "#0000ff"
	//	}
	//	color := z*(0x0000ff-0xff0000)/(math.Sqrt(2)*(xyrange*0.5)) + 0xff0000
	//	ret := fmt.Sprintf("#%x", int(color))
	return ret
}
