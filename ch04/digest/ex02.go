package digest

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
)

type shaType int

const (
	SHA256 shaType = iota
	SHA384
	SHA512
)

var SHAFlag shaType = SHA256

// MakeDigestはSHA256/SHA384/SHA512ハッシュを作成して返します。
// フラグでSHA256/SHA384/SHA512を切り替えます。
// デフォルトはSHA256です。
func MakeDigest(s string) string {
	var ret string

	switch SHAFlag {
	case SHA384:
		c := sha512.Sum384([]byte(s))
		ret = fmt.Sprintf("%x", c)
	case SHA512:
		c := sha512.Sum512([]byte(s))
		ret = fmt.Sprintf("%x", c)
	default:
		c := sha256.Sum256([]byte(s))
		ret = fmt.Sprintf("%x", c)
	}

	return ret
}
