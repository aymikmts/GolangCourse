package charcount

import (
	"bufio"
)

// WordFreqは単語の出力頻度を報告します。
func WordFreq(input *bufio.Scanner) *map[string]float64 {
	freq := make(map[string]float64) // 単語ごとの出力頻度
	counts := make(map[string]int)   // 単語ごとのカウント数
	num := 0                         // 単語総数

	input.Split(bufio.ScanWords)
	for input.Scan() {
		word := input.Text()
		counts[word]++
		num++
	}

	// 単語ごとの出現頻度を計算
	for word, count := range counts {
		freq[word] = float64(count) / float64(num) * 100.0
	}

	return &freq
}
