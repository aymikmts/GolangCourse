package dedup

// Dedupは[]stringスライス内で隣接している重複を除去します。
func Dedup(s []string) (ret []string) {

	for i := 0; i < len(s); i++ {
		if i == 0 {
			ret = append(ret, s[0])
			continue
		}

		if s[i] == s[i-1] {
			continue
		}

		ret = append(ret, s[i])
	}

	return
}
