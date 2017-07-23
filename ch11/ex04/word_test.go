package word

import (
	"math/rand"
	"testing"
	"time"
)

// randomPalindromeは、擬似乱数生成器rngから長さと内容が計算された
// 回文を返します。
func randomPalindrome(rng *rand.Rand) string {
	n := rng.Intn(25) // 24までのランダムな長さ

	// コンマと空白の挿入位置を決める
	var comma int
	for {
		if n == 0 {
			break
		}
		comma = rng.Intn(25)
		if comma < n {
			break
		}
	}
	var space int
	for {
		if n == 0 {
			break
		}
		space = rng.Intn(25)
		if space < n {
			break
		}
	}

	runes := make([]rune, n)
	for i := 0; i < (n+1)/2; i++ {
		r := rune(rng.Intn(0x1000)) // '\u0999'までのランダムなルーン
		runes[i] = r
		runes[n-1-i] = r
	}

	var ret string
	if comma < space {
		ret = string(runes[:comma]) + "," + string(runes[comma:space]) + " " + string(runes[space:])
	} else {
		ret = string(runes[:space]) + " " + string(runes[space:comma]) + "," + string(runes[comma:])
	}
	return ret
}

func TestRandomPalindromes(t *testing.T) {
	seed := time.Now().UTC().UnixNano()
	t.Logf("Random seed: %d", seed)
	rng := rand.New(rand.NewSource(seed))

	for i := 0; i < 1000; i++ {
		p := randomPalindrome(rng)
		if !IsPalindrome(p) {
			t.Errorf("IsPalindome(%q) = false", p)
		}
	}
}
