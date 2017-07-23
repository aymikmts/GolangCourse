package word

import (
	"math/rand"
	"testing"
	"time"
	"unicode"
)

// randomNotPalindromeは、擬似乱数生成器rngから長さと内容が計算された
// 回文でない文字列を返します。
func randomNotPalindrome(rng *rand.Rand) string {
	var n int
	// 文字列長が1より大きいものにする
	for {
		n = rng.Intn(25) // 24までのランダムな長さ
		if n > 2 {       // コンマを入れること場合を想定して、3文字以上になるようにする
			break
		}
	}

	// IsLetterでない文字列を挿入する位置を取得
	comma := rng.Intn(25)

	runes := make([]rune, n)
	for i := 0; i < n; i++ {
		// commaの位置に、","を挿入する
		if i == comma {
			runes[i] = ','
			continue
		}
		r := rune(rng.Intn(0x1000)) // '\u0999'までのランダムなルーン
		// rがrunes[n-1-i]と同じだったら、rを再取得する
		for {
			if unicode.ToLower(r) == unicode.ToLower(runes[n-1-i]) {
				r = rune(rng.Intn(0x1000))
			} else if !unicode.IsLetter(r) {
				r = rune(rng.Intn(0x1000))
			} else {
				break
			}
		}
		runes[i] = r
	}
	return string(runes)
}

func TestRandomNotPalindromes(t *testing.T) {
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	for i := 0; i < 1000; i++ {
		p := randomNotPalindrome(rng)
		if IsPalindrome(p) {
			t.Errorf("IsPalindome(%q) = true", p)
		}
	}
}
